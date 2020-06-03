package nsqout

import (
	"time"

	"github.com/elastic/beats/v7/libbeat/outputs/codec"
)

type nsqConfig struct {
	Nsqd         string        `config:"nsqd"`
	Topic        string        `config:"topic"`
	BulkMaxSize  int           `config:"bulk_max_size"`
	MaxRetries   int           `config:"max_retries"`
	WriteTimeout time.Duration `config:"write_timeout"`
	DialTimeout  time.Duration `config:"dial_timeout"`

	Codec codec.Config `config:"codec"`
}

var (
	defaultConfig = nsqConfig{
		WriteTimeout: 3 * time.Second,
		DialTimeout:  4 * time.Second,
		BulkMaxSize:  256,
		MaxRetries:   3,
		Nsqd:         "127.0.0.1:4150",
		Topic:        "nsqbeat",
	}
)
