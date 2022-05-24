package handler

import (
	"context"
	"post_model_predict/idl/post_model_predict"
	"post_model_predict/model"
	"post_model_predict/mq"
	"post_model_predict/util"
)

type PredictAysncHandler struct {
	Request  *post_model_predict.PredictAysncRequest
	Response *post_model_predict.PredictAysncResponse
	Ctx      context.Context
}

func NewPredictAysncHandler(Ctx context.Context, Request *post_model_predict.PredictAysncRequest) *PredictAysncHandler {
	return &PredictAysncHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_predict.PredictAysncResponse{
			Message: "success",
			Status:  0,
			Data:    &post_model_predict.PredictAysncData{},
		},
	}
}
func (h *PredictAysncHandler) Run() {
	task := h.CreatePredictTask()
	mq.SendPredictTask(h.Ctx, task)
	h.Response.Data.TaskId = task.TaskId
}

func (h *PredictAysncHandler) CheckParam() error {

	return nil
}

func (h *PredictAysncHandler) CreatePredictTask() *model.PredictTask {
	task := &model.PredictTask{
		Input:        *ConvertInputFromRpcToTorch(h.Request.Data),
		ModelName:    h.Request.ModelName,
		ModelVersion: h.Request.ModelVersion,
		TaskId:       util.GetUUID(),
	}
	return task
}
