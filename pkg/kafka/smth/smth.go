package smth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaClient struct {
	sarama.Client
}

func NewKafkaClient(kafkaHost, kafkaPort string) *KafkaClient {
	client, err := sarama.NewClient([]string{
		fmt.Sprintf("%s:%s", kafkaHost, kafkaPort),
	}, configureSarama())
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaClient{
		client,
	}
}

func (c *KafkaClient) Consumer(topics []string) error {
	consumer := c.getConsumer()

	return consumer.Consume(
		context.Background(),
		topics,
		ExampleConsumerGroupHandler{},
	)
}

func (c *KafkaClient) getConsumer() sarama.ConsumerGroup {
	consumer, err := sarama.NewConsumerGroupFromClient("group id", c.Client)
	if err != nil {
		panic(err)
	}
	return consumer
}

func configureSarama() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_0_0_0
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Producer.Timeout = 5 * time.Second
	return conf
}
