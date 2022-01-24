package main

import (
	"fmt"
	"wing.com/magic-golang/kafka/consumer"
	"wing.com/magic-golang/kafka/producer"
)

// kafka - 示例

func main() {
	// 1. 初始化producer
	err := producer.Init([]string{"127.0.0.1:9092"}, "testTopic")
	if err != nil {
		fmt.Printf("Init kafka producer failed, error is %v.\n", err)
		return
	}
	fmt.Println("Init kafka producer succeed!")

	// 2. 初始化consumer
	err = consumer.Init([]string{"127.0.0.1:9092"}, "testTopic")
	if err != nil {
		fmt.Printf("Init kafka consumer failed, error is %v.\n", err)
		return
	}
	fmt.Println("Init kafka consumer succeed!")
}
