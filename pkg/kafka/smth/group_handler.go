package smth

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler interface {
	// Setup is run at the beginning of a new session, before ConsumeClaim.
	Setup(sarama.ConsumerGroupSession) error

	// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
	// but before the offsets are committed for the very last time.
	Cleanup(sarama.ConsumerGroupSession) error

	// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
	// Once the Messages() channel is closed, the Handler must finish its processing
	// loop and exit.
	ConsumeClaim(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error
}

type ExampleConsumerGroupHandler struct{}

func (ExampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ExampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ExampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println("consume claim has been started")
	for {
		select {
		case msg := <-claim.Messages():
			fmt.Printf("Received message: %s, topic: %s, partition %v\n",
				string(msg.Value), msg.Topic, msg.Partition)
		}
	}

	// for {
	// 	select {
	// 	case msg := <-partitionConsumer.Messages():
	// 		println("Received message: ", string(msg.Value), "topic: ", topic, "partition: ", msg.Partition, "offset: ", msg.Offset)
	// 	case err = <-partitionConsumer.Errors():
	// 		println("Error: ", err.Error())
	// 		return err
	// 	}
	// }

	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}
