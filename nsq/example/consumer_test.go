package example

import (
	"fmt"
	gnsq "github.com/nsqio/go-nsq"
	"github.com/phuhao00/broker/nsq"
	"testing"
)

func TestConsumer(t *testing.T) {
	consumerClient := nsq.NewConsumerClient("127.0.0.1", "test", "tch", "4150", nil)
	ch := &ConsumerHandle{}
	ch.Register(0, func(message *gnsq.Message) error {
		fmt.Println("hello nsq")
		return nil
	})
	consumerClient.AddHandle(ch)
	select {}
}
