package test

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
	"time"
)

var topic = "canal-topic"

func TestKafkaProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder("hello kafka NB")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.238.128:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

// 192.168.124.33

func TestKafka(t *testing.T) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Retention = time.Second
	config.Version = sarama.V3_2_0_0
	consumer, err := sarama.NewConsumer([]string{"192.168.238.128:9092"}, config)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		for msg := range pc.Messages() {
			fmt.Println(msg.Topic)
			fmt.Printf("%v", string(msg.Value))
			fmt.Println(msg.Partition)

			//fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
		}
		pc.Close()
		// 异步从每个分区消费信息
		/*go func(sarama.PartitionConsumer) {
			fmt.Println("sarama partitionConsumer")
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)*/
	}

}
