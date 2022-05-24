package mq

import (
	"context"
	"encoding/json"
	"post_model_predict/model"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func SendPredictTask(ctx context.Context, task *model.PredictTask) error {
	taskJson, err := json.Marshal(task)
	if err != nil {
		return err
	}
	msg := &primitive.Message{
		Topic: predictTopic,
		Body:  taskJson,
	}
	msg.WithTag("predict_v_1")
	// err = Producer.SendAsync(ctx, callBack, msg)
	_, err = Producer.SendSync(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
