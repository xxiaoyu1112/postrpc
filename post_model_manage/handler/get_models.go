package handler

import (
	"context"
	"log"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/model"
	"post_model_manage/rpc"
)

type GetModelsHandler struct {
	Request  *post_model_manage.GetModelsRequest
	Response *post_model_manage.GetModelsResponse
	Ctx      context.Context
}

func NewGetModelsHandler(Ctx context.Context, Request *post_model_manage.GetModelsRequest) *GetModelsHandler {
	return &GetModelsHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_manage.GetModelsResponse{
			Message: "success",
			Status:  0,
			Data:    &post_model_manage.GetModelsData{},
		},
	}
}
func (h *GetModelsHandler) Run() {
	ctx := h.Ctx
	torchModelIds, err := rpc.GetModelList(ctx)
	if err != nil {
		log.Printf("[Error] call rpc GetModelList, err: %v", err)
		// Error code now all define to 1
		h.Response.Message = err.Error()
		h.Response.Status = int32(1)
		return
	}
	h.Response.Data.Data = h.ConvertModelIdFromTorchToRpc(torchModelIds)
}

func (h *GetModelsHandler) CheckParam() error {

	return nil
}

func (h *GetModelsHandler) ConvertModelIdFromTorchToRpc(torchModelId *model.ModelListResponse) []*post_model_manage.ModelIdentity {
	var rpcModelIds []*post_model_manage.ModelIdentity
	for _, torchModelId := range torchModelId.Models {
		rpcModelIds = append(rpcModelIds, &post_model_manage.ModelIdentity{
			ModelName:    torchModelId.ModelName,
			ModelVersion: torchModelId.ModelName,
		})
	}
	return rpcModelIds
}
