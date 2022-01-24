package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka consumer

var (
	consumer sarama.Consumer
)

func Init(address []string, topic string) (err error) {
	consumer, err = sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	go fetchMSgFromProducer(topic)

	return
}

func fetchMSgFromProducer(topic string) (err error) {
	partitionList, err = consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
		// select {}
	}
	return
}

//func main() {
//	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
//	if err != nil {
//		fmt.Printf("fail to start consumer, err:%v\n", err)
//		return
//	}
//	partitionList, err := consumer.Partitions("testTopic") // 根据topic取到所有的分区
//	if err != nil {
//		fmt.Printf("fail to get list of partition:err%v\n", err)
//		return
//	}
//	fmt.Println(partitionList)
//	for partition := range partitionList { // 遍历所有的分区
//		// 针对每个分区创建一个对应的分区消费者
//		pc, err := consumer.ConsumePartition("testTopic", int32(partition), sarama.OffsetNewest)
//		if err != nil {
//			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
//			return
//		}
//		defer pc.AsyncClose()
//		// 异步从每个分区消费信息
//		go func(sarama.PartitionConsumer) {
//			for msg := range pc.Messages() {
//				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
//			}
//		}(pc)
//		select {}
//	}
//}
