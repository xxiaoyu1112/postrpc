package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"post_model_predict/model"
	"time"

	Redis "github.com/go-redis/redis/v8"
)

const predictTaskPrefix = "predict:"

func GetPredictTaskResult(ctx context.Context, taskId string) (*model.PredictTaskResult, error) {
	val, err := rdb.Get(ctx, predictTaskPrefix+taskId).Result()
	if err == Redis.Nil {
		return nil, fmt.Errorf("task not exist error, task_id :%v ", taskId)
	} else if err != nil {
		return nil, err
	}
	result := model.PredictTaskResult{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, fmt.Errorf("task unmarshal error,task_id: %v, err: %v", taskId, err)
	}
	return &result, nil
}

func SetPredictTaskResult(ctx context.Context, res *model.PredictTaskResult) error {
	resMsg, err := json.Marshal(res)
	if err != nil {
		return err
	}
	key := predictTaskPrefix + res.TaskId
	val := string(resMsg)
	err = rdb.Set(ctx, key, val, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}
