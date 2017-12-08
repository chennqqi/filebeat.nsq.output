package nsqout

import (
	"time"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/codec"
	"github.com/elastic/beats/libbeat/outputs/outil"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/nsqio/go-nsq"
)

type publishFn func(
	keys outil.Selector,
	data []publisher.Event,
) ([]publisher.Event, error)

type client struct {
	outputs.NetworkClient

	stats outputs.Observer
	codec codec.Codec
	index string

	//for nsq
	nsqd     string
	topic    string
	producer *nsq.Producer
	config   *nsq.Config
}

func newClient(
	stats outputs.Observer,
	codec codec.Codec,
	index string,
	nsqd string,
	topic string,
	writeTimeout time.Duration,
	dialTimeout time.Duration,
) *client {
	cfg := nsq.NewConfig()
	cfg.WriteTimeout = writeTimeout
	cfg.DialTimeout = dialTimeout
	return &client{
		codec:  codec,
		stats:  stats,
		index:  index,
		nsqd:   nsqd,
		config: cfg,
		topic:  topic,
	}
}

func (c *client) Connect() error {
	debugf("connect to %s", c.nsqd)

	producer, err := nsq.NewProducer(c.nsqd, c.config)
	if err != nil {
		logp.Err("[main:NsqForward.Open] NewProducer error ", err)
		return err
	}
	//TODO: set logger
	//producer.SetLogger(nullLogger, LogLevelInfo)
	c.producer = producer
	return err
}

func (c *client) Publish(batch publisher.Batch) error {
	if c == nil {
		panic("no client")
	}
	if batch == nil {
		panic("no batch")
	}

	events := batch.Events()
	c.stats.NewBatch(len(events))

	st := c.stats

	//build message failed
	msgs, err := c.buildNsqMessages(events)
	dropped := len(events) - len(msgs)
	if err != nil {
		//	st.Dropped(dropped)
		//	st.Acked(len(events) - dropped)
		logp.Err("[main:nsq] c.buildNsqMessages ", err)
		c.stats.Failed(len(events))
		batch.RetryEvents(events)
		return err
	}

	//nsq send failed do retry...
	err = c.producer.MultiPublish(c.topic, msgs)
	if err != nil {
		logp.Err("[main:nsq] producer.MultiPublish ", err)
		c.stats.Failed(len(events))
		batch.RetryEvents(events)
		return err
	}
	batch.ACK()

	st.Dropped(dropped)
	st.Acked(len(msgs))
	return err
}

func (c *client) buildNsqMessages(events []publisher.Event) ([][]byte, error) {
	length := len(events)
	msgs := make([][]byte, length)
	var count int

	var err error
	for idx := 0; idx < length; idx++ {
		event := events[idx].Content
		serializedEvent, err := c.codec.Encode(c.index, &event)
		if err != nil {
			logp.Err("[main:nsq] c.codec.Encode ", err)
			err = err
		} else {
			msgs[count] = serializedEvent
			count++
		}
	}
	return msgs[:count], err
}
