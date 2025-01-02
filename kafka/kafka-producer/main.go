package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		panic(err.Error())
	}
	defer producer.Close()

	fmt.Println("kafka is up and running", producer)

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

	topic := "user_create"
	message := "Hello PrimeSoft"
	err = producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic: &topic, Partition: kafka.PartitionAny,
	}, Value: []byte(message),
	}, nil)
	if err != nil {
		fmt.Println("could not publish a message, due to ", err.Error())
	} else {
		fmt.Println("message successfully sent")
	}
	producer.Flush(15000)
}
