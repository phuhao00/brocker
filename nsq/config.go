package nsq

import (
	"github.com/nsqio/go-nsq"
	"time"
)

type ConsumerConfig struct {
	Topic        string
	Channel      string
	Address      string
	Lookup       []string
	MaxInFlight  int
	Identify     nsq.IdentifyResponse
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DrainTimeout time.Duration
}

const (
	DefaultUserAgent       = ""
	DefaultMaxConcurrency  = 1
	DefaultMaxInFlight     = 1
	DefaultDialTimeout     = 5 * time.Second
	DefaultReadTimeout     = 1 * time.Minute
	DefaultWriteTimeout    = 10 * time.Second
	DefaultLookupTimeout   = 10 * time.Second
	DefaultMaxRetryTimeout = 10 * time.Second
	DefaultMinRetryTimeout = 10 * time.Millisecond
	DefaultDrainTimeout    = 10 * time.Second

	NoTimeout = time.Duration(0)
)

func (c *ConsumerConfig) defaults() {
	if c.MaxInFlight == 0 {
		c.MaxInFlight = DefaultMaxInFlight
	}

	if c.DialTimeout == 0 {
		c.DialTimeout = DefaultDialTimeout
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = DefaultReadTimeout
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = DefaultWriteTimeout
	}
	if c.DrainTimeout == 0 {
		c.DrainTimeout = DefaultDrainTimeout
	}
}
