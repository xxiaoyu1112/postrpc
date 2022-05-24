package mq

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Producer rocketmq.Producer

const mqAddress = "211.71.76.189:9876"
const predictTopic = "post_model_predict_task"

func init() {
	p, _ := rocketmq.NewProducer(
		// 设置  nameSrvAddr
		// nameSrvAddr 是 Topic 路由注册中心
		producer.WithNameServer([]string{mqAddress}),
		// 指定发送失败时的重试时间
		producer.WithRetry(2),
		// 设置 Group
		producer.WithGroupName("task_predict_Group"),
	)
	// 开始连接
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		panic(err)
	}
	Producer = p
}

var callBack = func(ctx context.Context, result *primitive.SendResult, err error) {
	if err != nil {
		log.Printf("[Error] send message err:%v ", err)
	}
}
