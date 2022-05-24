package mysql

import (
	"context"
	"errors"
	"log"
	"post_data_manage/model"

	"gorm.io/gorm"
)

func CreateTask(ctx context.Context, task *model.PostTask) error {
	err := DB.Create(task).Error
	if err != nil {
		log.Printf("[CreateTask] error query mysql, err: %v", err)
		return err
	}
	return nil
}

func GetTaskById(ctx context.Context, taskId string) (*model.PostTask, error) {
	task := &model.PostTask{}
	err := DB.Where("post_task_id = ?", taskId).First(task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Printf("[CreateTask] error query mysql, err: %v", err)
		return nil, err
	}
	return task, nil
}
