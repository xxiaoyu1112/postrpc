package model

type ModelListResponse struct {
	Models []Models `json:"models"`
}

type Models struct {
	ModelName string `json:"modelName"`
	ModelUrl  string `json:"modelUrl"`
}
