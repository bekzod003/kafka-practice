package main

import (
	"time"

	"github.com/bekzod003/kafka-practice/config"
	"github.com/bekzod003/kafka-practice/pkg/kafka/producer"
)

func main() {
	cfg := config.NewConfig("localhost", "29092", "test")

	producer := producer.NewProducer(
		cfg.KafkaHost,
		cfg.KafkaPort,
	)

	const topic = "lol"
	go func() {
		for {
			producer.Produce(topic, "Hello, world!")
			producer.Produce("bekzod", "Mbappe")
			time.Sleep(time.Second * 5)
		}
	}()
	for {

	}
	//
	// consumer1 := consumer.NewConsumer(cfg.KafkaHost, cfg.KafkaPort)
	// go func() {
	// 	if err := consumer1.Consume(topic); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	//
	// if err := consumer.NewConsumer(cfg.KafkaHost, cfg.KafkaPort).Consume("bekzod"); err != nil {
	// 	log.Fatal(err)
	// }

}
