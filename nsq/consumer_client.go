package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	_ "github.com/nsqio/go-nsq"
	"log"
	"net"
)

type ConsumerClient struct {
	q                *nsq.Consumer
	messagesSent     int
	messagesReceived int
	messagesFailed   int
	NSQDPort         string
}

func NewConsumerClient(addr, topicName, channel, NSQDPort string, cb func(c *nsq.Config)) *ConsumerClient {
	config := nsq.NewConfig()
	config.LocalAddr, _ = net.ResolveTCPAddr("tcp", addr+":0")
	//config.DefaultRequeueDelay = 0
	//config.MaxBackoffDuration = time.Millisecond * 50
	if cb != nil {
		cb(config)
	}
	if config.Deflate {
		topicName = topicName + "_deflate"
	} else if config.Snappy {
		topicName = topicName + "_snappy"
	}
	if config.TlsV1 {
		topicName = topicName + "_tls"
	}
	q, _ := nsq.NewConsumer(topicName, channel, config)
	q.SetLogger(log.Default(), nsq.LogLevelDebug)
	c := &ConsumerClient{q: q, NSQDPort: NSQDPort}

	return c
}

func (c *ConsumerClient) AddHandle(handler nsq.Handler) {
	c.q.AddHandler(handler)
	err := c.q.ConnectToNSQD("127.0.0.1:" + c.NSQDPort)
	if err != nil {
		fmt.Println(err)
	}
}
