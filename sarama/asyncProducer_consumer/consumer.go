package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka 브로커 리스트
	brokers := []string{"localhost:9092"}
	// 구독할 Kafka 토픽
	topic := "test_topic"
	// Consumer Group 이름
	groupID := "example-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0 // Kafka 버전에 맞게 설정
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 처음부터 메시지를 소비하도록 설정

	// Consumer Group 생성
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer func() {
		if err := consumerGroup.Close(); err != nil {
			log.Fatalf("Error closing consumer group: %v", err)
		}
	}()

	// Consumer Group 핸들러 생성
	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx := context.Background()

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, &consumer); err != nil {
				log.Fatalf("Error consuming messages: %v", err)
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Consumer가 준비될 때까지 대기
	log.Println("Sarama consumer group up and running...")

	// 프로그램 종료 대기
	select {}
}

// Consumer 구조체 정의
type Consumer struct {
	ready chan bool
}

// Setup은 컨슈머 세션이 시작될 때 호출
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

// Cleanup은 컨슈머 세션이 종료될 때 호출
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim은 메시지 소비를 처리
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: topic = %s, partition = %d, offset = %d, value = %s\n",
			message.Topic, message.Partition, message.Offset, string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}
