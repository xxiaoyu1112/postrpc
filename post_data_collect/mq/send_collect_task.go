package mq

import (
	"context"
	"encoding/json"
	"post_data_collect/model"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func SendCollectTask(ctx context.Context, task *model.CollectTask) error {
	taskJson, err := json.Marshal(task)
	if err != nil {
		return err
	}
	msg := &primitive.Message{
		Topic: predictTopic,
		Body:  taskJson,
	}
	msg.WithTag("collect_v1_message")
	// err = Producer.SendAsync(ctx, callBack, msg)
	_, err = Producer.SendSync(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
