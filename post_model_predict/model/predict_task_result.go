package model

// Redis中存在的结果结构体
type PredictTaskResult struct {
	TaskId     string  `json:"task_id"`
	TaskStatus int32   `json:"task_status"`
	TaskResult []int32 `json:"task_result"`
}
