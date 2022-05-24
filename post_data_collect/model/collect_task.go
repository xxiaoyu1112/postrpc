package model

type CollectTask struct {
	TaskId      string `json:"uuid"`
	TaskContent string `json:"task_content"`
}
