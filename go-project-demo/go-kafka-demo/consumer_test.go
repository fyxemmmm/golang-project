package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"sync"
	"testing"
)

type consumerGroupHandler struct {
	name string
}

// 每一个Topic的分区只能被一个消费组中的一个消费者所消费。一个消费者可以同时消费多个topic
func TestConsumer(t *testing.T) {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewClient([]string{"xx.fyxemmmm.cn:9092"}, config)
	
	
	defer client.Close()
	if err != nil {
		panic(err)
	}

/*	client2, err := sarama.NewClient([]string{"txy.fyxemmmm.cn:9092"}, config)
	if err != nil {
		panic(err)
	}
	client2Group, err := sarama.NewConsumerGroupFromClient("c3", client2)
	if err != nil {
		panic(err)
	}
	defer client2Group.Close()*/


	group1, err := sarama.NewConsumerGroupFromClient("c1", client)
	if err != nil {
		panic(err)
	}
	group2, err := sarama.NewConsumerGroupFromClient("c2", client)
	if err != nil {
		panic(err)
	}
	group3, err := sarama.NewConsumerGroupFromClient("c3", client)
	if err != nil {
		panic(err)
	}

	defer group1.Close()
	defer group2.Close()
	defer group3.Close()
	wg.Add(3)
	go consume(group1, &wg, "c1")
	go consume(group2, &wg, "c2")
	go consume(group3, &wg, "c3")
	
	//go consume(client2Group, &wg, "c3")
	wg.Wait()
	
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}

func consume(group sarama.ConsumerGroup, wg *sync.WaitGroup, name string) {
	fmt.Println(name + "start")
	wg.Done()
	ctx := context.Background()
	for {
		//topic := []string{"tiantian_topic1","tiantian_topic2"} 可以消费多个topic
		topics := []string{"fyx_topic"}
		handler := consumerGroupHandler{name: name}
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

/*
func handleErrors(group *sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Done()
	for err := range (*group).Errors() {
		fmt.Println("ERROR", err)
	}
}
*/