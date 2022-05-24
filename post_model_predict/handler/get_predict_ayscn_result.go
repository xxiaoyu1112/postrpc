package handler

import (
	"context"
	"post_model_predict/idl/post_model_predict"
	"post_model_predict/redis"
)

type GetPredictAysncResultHandler struct {
	Request  *post_model_predict.GetPredictAysncResultRequest
	Response *post_model_predict.GetPredictAysncResultResponse
	Ctx      context.Context
}

func NewGetPredictAysncResultHandler(Ctx context.Context, Request *post_model_predict.GetPredictAysncResultRequest) *GetPredictAysncResultHandler {
	return &GetPredictAysncResultHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_predict.GetPredictAysncResultResponse{
			Message: "success",
			Status:  0,
			Data:    &post_model_predict.PredictAysncResult{},
		},
	}
}
func (h *GetPredictAysncResultHandler) Run() {
	taskId := h.Request.TaskId
	result, err := redis.GetPredictTaskResult(h.Ctx, taskId)
	if err != nil {
		h.Response.Status = 1
		h.Response.Message = err.Error()
		return
	}
	h.Response.Data.Status = result.TaskStatus
	h.Response.Data.Order = result.TaskResult
}

func (h *GetPredictAysncResultHandler) CheckParam() error {

	return nil
}
