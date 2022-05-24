package model

import "time"

type PostTask struct {
	PostTaskId     string    `gorm:"column:post_task_id;" json:"postTaskId"`
	PostTaskDate   time.Time `gorm:"column:post_task_date;" json:"postTaskDate"`
	PostTaskStatus int32     `gorm:"column:post_task_status;" json:"postTaskStatus"`
}

func (t *PostTask) TableName() string {
	return "post_task"
}
