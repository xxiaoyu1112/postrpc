package handler

import (
	"context"
	"log"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/rpc"
)

type RemoveModelHandler struct {
	Request  *post_model_manage.RemoveModelRequest
	Response *post_model_manage.RemoveModelResponse
	Ctx      context.Context
}

func NewRemoveModelHandler(Ctx context.Context, Request *post_model_manage.RemoveModelRequest) *RemoveModelHandler {
	return &RemoveModelHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_manage.RemoveModelResponse{
			Message: "success",
			Status:  0,
		},
	}
}
func (h *RemoveModelHandler) Run() {
	ctx := h.Ctx
	modelId := h.Request.ModelIdentity
	err := rpc.RemoveModel(ctx, modelId.ModelName, modelId.ModelVersion)
	if err != nil {
		log.Printf("[Error] call rpc RemoveModel, err: %v", err)
		// Error code now all define to 1
		h.Response.Message = err.Error()
		h.Response.Status = int32(1)
		return
	}
}

func (h *RemoveModelHandler) CheckParam() error {

	return nil
}
