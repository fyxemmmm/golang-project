package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
)

// 同步客户端
func TestSyncProducer(t *testing.T) {
	config := sarama.NewConfig()                                              //实例化个sarama的Config
	config.Producer.Return.Successes = true                                   //是否开启消息发送成功后通知 successes channel
	config.Producer.Partitioner = sarama.NewRandomPartitioner                 //随机分区器
	client, err := sarama.NewClient([]string{"xx.fyxemmmm.cn:9092"}, config) //初始化客户端
	defer client.Close()
	if err != nil {
		panic(err)
	}
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{Topic: "fyx_topic", Key: nil, Value: sarama.StringEncoder("yyy")})
	if err != nil {
		log.Fatalf("unable to produce message: %q", err)
	}
	fmt.Println("partition", partition)
	fmt.Println("offset", offset)
}

// 异步客户端  ->  可以把数据线缓存成batch, 内部积攒了batchs -> 一起发送给broker
func TestAsyncProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // true就变成了同步方式  注意此时select需要去掉default
	
	client, err := sarama.NewClient([]string{"txy.fyxemmmm.cn:9092"}, config)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("unable to create kafka producer: %q", err)
	}
	defer producer.Close()

	text := "hello world"
	producer.Input() <- &sarama.ProducerMessage{Topic: "fyx_topic", Key: nil, Value: sarama.StringEncoder(text)}
	// wait response
	select {
	case msg := <-producer.Successes():  // 同步模式这个必须要取走, 不然success的chan会被写满
	  log.Printf("Produced message successes: [%s]\n",msg.Value)
	case err := <-producer.Errors():
		log.Println("Produced message failure: ", err)
	default:
		log.Println("Produced message default")
	}
	
	// 配置config.Producer.Return.Successes = true和操作<-producer.Successes()必须配套使用；
	// 配置成true，那么就要去读取Successes，如果配置成false，则不能去读取Successes。
}