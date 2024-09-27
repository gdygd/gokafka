package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

var config *sarama.Config = nil
var client sarama.Client = nil
var consumer sarama.Consumer = nil
var partitionConsumer sarama.PartitionConsumer = nil

func connectK() {
	var err error = nil

	// Kafka 브로커 주소
	brokers := []string{"localhost:9092"}

	// Sarama 설정
	config = sarama.NewConfig()
	config.Producer.Return.Successes = true

	log.Println("#1 connectK")
	client, err = sarama.NewClient(brokers, config)
	if err != nil {
		log.Printf("Kafka 클라이언트 생성 실패: %v\n", err)
		return
	}
	log.Println("#2 connectK")

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("Kafka 컨슈머 생성 실패: %v", err)
	}
	log.Println("#3 connectK")

	// 특정 토픽의 파티션 컨슈머 생성
	partitionConsumer, err = consumer.ConsumePartition("test_topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("파티션 컨슈머 생성 실패: %v", err)
	}

}

func clearClient() {
	if client != nil && !client.Closed() {
		client.Close()
		log.Println("clearClient")
	}
	client = nil
}
func clearConsumer() {
	if consumer != nil {
		consumer.Close()
		log.Println("clearConsumer")
	}
	consumer = nil
}

func clearPartitionConsumer() {
	if partitionConsumer != nil {
		partitionConsumer.Close()
		log.Println("clearPartitionConsumer")
	}
	partitionConsumer = nil
}

func main() {

	connectK()
	defer func() {
		clearClient()
		clearConsumer()
		clearPartitionConsumer()
	}()

	// 메시지 수신
	// for message := range partitionConsumer.Messages() {
	// 	log.Printf("메시지 수신: 파티션=%d, 오프셋=%d, 키=%s, 값=%s\n", message.Partition, message.Offset, string(message.Key), string(message.Value))
	// }

	for {
		if partitionConsumer != nil {
			select {
			case message := <-partitionConsumer.Messages():
				log.Printf("메시지 수신: 파티션=%d, 오프셋=%d, 키=%s, 값=%s\n", message.Partition, message.Offset, string(message.Key), string(message.Value))
			default:
				log.Println("not receive message")

				// // close and connect
				clearClient()
				clearConsumer()
				clearPartitionConsumer()
				// connectK()
			}
		}
		if client == nil {
			connectK()
		}
		fmt.Println("read msg.....")

		time.Sleep(time.Second * 1)
	}

}
