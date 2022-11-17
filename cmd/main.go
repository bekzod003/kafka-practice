package main

import (
	"log"
	"time"

	"github.com/bekzod003/kafka-practice/config"
	"github.com/bekzod003/kafka-practice/pkg/kafka/consumer"
	"github.com/bekzod003/kafka-practice/pkg/kafka/producer"
)

func main() {
	cfg := config.NewConfig("localhost", "9092", "test")

	producer := producer.NewProducer(
		cfg.KafkaHost,
		cfg.KafkaPort,
	)

	const topic = "lol"
	go func() {
		for {
			producer.Produce(topic, "Hello, world!")
			time.Sleep(time.Second * 5)
		}
	}()

	consumer := consumer.NewConsumer(cfg.KafkaHost, cfg.KafkaPort)
	if err := consumer.Consume(topic); err != nil {
		log.Fatal(err)
	}
}
