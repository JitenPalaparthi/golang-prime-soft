package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "common_gouup-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err.Error())
	}
	defer consumer.Close()

	topic := "user_create"

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Println("error in subscribing to the topics", err.Error())
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(msg.Value))
	}

}

// 1 Pt -- 1 C
// 2 Pt -- 1 C
// 2 Pt 10 M  -- 2 C
// 2 Pt -- 3 C (waste of resources)
// 3 Pt -- 2 C
// 3 Pt -- 3 C
// 2 Pt -- 2 C (2 CGs)
