package producer

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(kafkaHost, kafkaPort string) *Producer {
	producer, err := sarama.NewSyncProducer([]string{
		fmt.Sprintf("%s:%s", kafkaHost, kafkaPort),
	}, configureSarama())

	if err != nil {
		log.Fatal(err)
	}

	return &Producer{
		producer: producer,
	}
}

func (p *Producer) Produce(topic string, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message is stored in topic %s", topic)
}

func configureSarama() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_0_0_0
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Producer.Timeout = 5 * time.Second
	return conf
}
