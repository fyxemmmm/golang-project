package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func TestConsumer2(t *testing.T) {
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"xx.fyxemmmm.cn:9092"}, nil)
	if err != nil {
		panic(err)
	}
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions("fyx_topic")
	if err != nil {
		panic(err)
	}
	fmt.Println(partitionList)

	// 随机的consumer group
	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition("fyx_topic", int32(partition), sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}