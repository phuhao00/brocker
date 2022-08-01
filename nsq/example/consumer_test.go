package example

import (
	"fmt"
	gnsq "github.com/nsqio/go-nsq"
	"github.com/phuhao00/broker/nsq"
	"testing"
)

func TestConsumer(t *testing.T) {
	conf := nsq.ConsumerConfig{
		Topic:   "test",
		Channel: "tch",
		Address: "127.0.0.1",
		Lookup:  []string{"127.0.0.1:4150"},
	}
	consumerClient := nsq.NewConsumerClient(conf, nil)
	ch := &ConsumerHandle{}
	ch.Register(0, func(message *gnsq.Message) error {
		fmt.Println("hello nsq")
		return nil
	})
	consumerClient.AddHandle(ch)
	select {}
}
