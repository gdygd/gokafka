package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	var produc sarama.SyncProducer
	var err error = nil
	//producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	produc, err = sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	//producer, err := sarama.NewSyncProducer([]string{"localhost:2181"}, nil)
	if err != nil {
		log.Fatal("Producer creation failed:", err)
	}
	defer produc.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("Hello Kafka"),
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			partition, offset, err := produc.SendMessage(msg)
			if err != nil {
				//log.Fatal("Message send failed:", err)
				log.Println("Message send failed:", err)
				produc.Close()
				produc, err = sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
			}
			log.Printf("Message sent to partition %d with offset %d\n", partition, offset)
			time.Sleep(time.Second * 1)
		}
	}()

	// partition, offset, err := producer.SendMessage(msg)
	// if err != nil {
	// 	log.Fatal("Message send failed:", err)
	// }

	// log.Printf("Message sent to partition %d with offset %d\n", partition, offset)

	<-signals
}
