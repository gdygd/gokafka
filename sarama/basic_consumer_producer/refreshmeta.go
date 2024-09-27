package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// Sarama 라이브러리를 사용하는 경우
	client, err := sarama.NewClient([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Kafka 클라이언트 생성 실패: %v", err)
	}

	// 토픽의 메타데이터 갱신
	err = client.RefreshMetadata("test_topic")
	if err != nil {
		log.Fatalf("메타데이터 갱신 실패: %v", err)
	}

}
