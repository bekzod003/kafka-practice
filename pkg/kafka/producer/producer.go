package producer

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.AsyncProducer
}

func NewProducer(kafkaHost, kafkaUrl string) *Producer {
	producer, err := sarama.NewAsyncProducer([]string{
		fmt.Sprintf("%s:%s", kafkaHost, kafkaUrl),
	}, configureSarama())

	if err != nil {
		log.Fatal(err)
	}

	return &Producer{
		producer: producer,
	}
}

func (p *Producer) Produce(topic string, message string) {
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	println("SUccessfully produced message: ", message, "topic: ", topic)
}

func configureSarama() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_0_0_0
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Producer.Timeout = 5 * time.Second
	return conf
}
