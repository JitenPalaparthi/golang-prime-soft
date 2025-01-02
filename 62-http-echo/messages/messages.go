package messages

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type UserMessage struct {
	Servers string
	Topic   string
	Value   chan []byte
	//Key     chan []byte
}

func NewUserMessage(servers, topic string) *UserMessage {
	um := new(UserMessage)
	um.Servers = servers
	um.Topic = topic
	um.Value = make(chan []byte)
	//	um.Key = make(chan []byte)
	return um
}

func (um *UserMessage) Init() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": um.Servers,
	})

	if err != nil {
		panic(err.Error())
	}

	go func() {
		for event := range producer.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("delivery failed,", ev.TopicPartition.Error.Error())
				} else {
					fmt.Println("message delivered to  partition:", ev.TopicPartition)
				}
			}
		}
	}()
	//	defer producer.Close()

	go func() {

	}()
	go func() {
		for msg := range um.Value {
			err = producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
				Topic: &um.Topic, Partition: kafka.PartitionAny,
			}, Value: msg,
			}, nil)
			if err != nil {
				fmt.Println("could not publish a message, due to ", err.Error())
			} else {
				fmt.Println("message successfully sent")
			}
		}
	}()
}
