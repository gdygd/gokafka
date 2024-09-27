package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

var config *sarama.Config = nil
var client sarama.Client = nil
var producer sarama.SyncProducer = nil

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

	producer, err = sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Printf("Kafka 프로듀서 생성 실패: %v\n", err)
	}
	log.Println("#3 connectK")
}

func sendMessage() bool {
	message := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// 메시지 전송
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Printf("메시지 전송 실패: %v\n", err)
		return false
	}

	log.Printf("메시지 전송 성공: 파티션=%d, 오프셋=%d\n", partition, offset)

	return true

}

func clearClient() {
	if client != nil && !client.Closed() {
		client.Close()
	}
	client = nil
}
func clearProducer() {
	if producer != nil {
		producer.Close()
	}
	producer = nil
}

func main() {

	connectK()
	defer func() {
		clearClient()
		clearProducer()
	}()

	for {
		var isok bool
		log.Println("#4 connectK")
		if client != nil && !client.Closed() {
			log.Println("#5 connectK")
			isok = sendMessage()
			log.Println("#6 connectK")
		}

		if !isok {
			log.Println("#7 connectK")
			clearClient()
			log.Println("#8 connectK")
			clearProducer()

			log.Println("#9 connectK")
			connectK()
			log.Println("#10 connectK")

		}

		log.Println("#11 connectK")
		time.Sleep(time.Second * 1)
	}

}
