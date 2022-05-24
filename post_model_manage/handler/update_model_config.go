package handler

import (
	"context"
	"log"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/rpc"
)

type UpdateModelConfigHandler struct {
	Request  *post_model_manage.UpdateModelConfigRequest
	Response *post_model_manage.UpdateModelConfigResponse
	Ctx      context.Context
}

func NewUpdateModelConfigHandler(Ctx context.Context,Request *post_model_manage.UpdateModelConfigRequest) *UpdateModelConfigHandler {
	return &UpdateModelConfigHandler{
		Request: Request,
		Ctx: Ctx,
		Response: &post_model_manage.UpdateModelConfigResponse{
			Message: "success",
			Status:  0,
		},
	}
}
func (h *UpdateModelConfigHandler) Run() {
	ctx := h.Ctx
	modelId := h.Request.ModelIdentity
	modelConfig := h.Request.ModelConfig
	err :=rpc.UpdateModelConfig(ctx,modelId.ModelName,modelId.ModelVersion,modelConfig.MaxWorkers,modelConfig.MinWorkers)
	if err != nil{
		log.Printf("[Error] call rpc UpdateModelConfig, err: %v", err)
		// Error code now all define to 1
		h.Response.Message = err.Error()
		h.Response.Status = int32(1)
		return
	}
}

func (h *UpdateModelConfigHandler) CheckParam() error {

	return nil
}
