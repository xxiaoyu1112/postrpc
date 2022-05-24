package model

type ModelState struct {
	ModelName       string `json:"modelName"`
	ModelVersion    string `json:"modelVersion"`
	ModelUrl        string `json:"modelUrl"`
	Runtime         string `json:"runtime"`
	MinWorkers      int    `json:"minWorkers"`
	MaxWorkers      int    `json:"maxWorkers"`
	BatchSize       int    `json:"batchSize"`
	MaxBatchDelay   int    `json:"maxBatchDelay"`
	LoadedAtStartup bool   `json:"loadedAtStartup"`
	Workers         []struct {
		Id          string `json:"id"`
		StartTime   string `json:"startTime"`
		Status      string `json:"status"`
		MemoryUsage int    `json:"memoryUsage"`
		Pid         int    `json:"pid"`
		Gpu         bool   `json:"gpu"`
		GpuUsage    string `json:"gpuUsage"`
	} `json:"workers"`
}
