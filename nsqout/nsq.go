package nsqout

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
	//	"github.com/elastic/beats/libbeat/outputs/outil"
	"github.com/elastic/beats/libbeat/outputs/codec"
)

var debugf = logp.MakeDebug("nsq")

func init() {
	outputs.RegisterType("nsq", makeNsq)
}

func makeNsq(
	_ outputs.IndexManager,
	beat beat.Info,
	observer outputs.Observer,
	cfg *common.Config,
) (outputs.Group, error) {
	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	codec, err := codec.CreateEncoder(beat, config.Codec)
	if err != nil {
		return outputs.Fail(err)
	}

	//	if !cfg.HasField("topic") {
	//		cfg.SetString("topic", -1, beat.Beat)
	//	}
	//	if err != nil {
	//		return outputs.Fail(err)
	//	}

	client := newClient(observer, codec, beat.Beat, config.Nsqd, config.Topic, config.WriteTimeout, config.DialTimeout)
	outputs.Success(config.BulkMaxSize, config.MaxRetries, client)
	return outputs.Success(config.BulkMaxSize, config.MaxRetries, client)
}
