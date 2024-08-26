package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

func main() {
	brokerList := []string{"localhost:9092"}
	topic := "test_topic"

	producer, err := sarama.NewAsyncProducer(brokerList, nil)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	defer producer.AsyncClose()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case success := <-producer.Successes():
				fmt.Printf("Message sent to topic %s, partition %d, offset %d\n",
					success.Topic, success.Partition, success.Offset)
			case err := <-producer.Errors():
				fmt.Printf("Failed to produce message: %v\n", err)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
		}
		producer.Input() <- msg
	}

	<-signals
}
