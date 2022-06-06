package consumer

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

func NewConsumer(kafkaHost, kafkaUrl string) *Consumer {
	consumer, err := sarama.NewConsumer([]string{
		fmt.Sprintf("%s:%s", kafkaHost, kafkaUrl),
	}, configureSarama())

	if err != nil {
		log.Fatal(err)
	}

	return &Consumer{
		consumer: consumer,
	}
}

func (c *Consumer) Consume(topic string) (err error) {
	partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, -1)
	if err != nil {
		return err
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			println("Received message: ", string(msg.Value), "topic: ", topic, "partition: ", msg.Partition, "offset: ", msg.Offset)
		case err = <-partitionConsumer.Errors():
			println("Error: ", err.Error())
			return err
		}
	}
}

func configureSarama() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_0_0_0
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Producer.Timeout = 5 * time.Second
	return conf
}
