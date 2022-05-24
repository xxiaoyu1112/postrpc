package handler

import (
	"context"
	"log"
	"post_model_predict/idl/post_model_predict"
	"post_model_predict/model"
	"post_model_predict/rpc"
)

type PredictHandler struct {
	Request  *post_model_predict.PredictRequest
	Response *post_model_predict.PredictResponse
	Ctx      context.Context
}

func NewPredictHandler(Ctx context.Context, Request *post_model_predict.PredictRequest) *PredictHandler {
	return &PredictHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_predict.PredictResponse{
			Message: "success",
			Status:  0,
			Data:    &post_model_predict.PredictResult{},
		},
	}
}
func (h *PredictHandler) Run() {
	ctx := h.Ctx
	torchInput := ConvertInputFromRpcToTorch(h.Request.Data)
	torchPredictResult, err := rpc.Predict(ctx, h.Request.ModelName, h.Request.ModelVersion, torchInput)
	if err != nil {
		log.Printf("[PredictHandler] error call rpc Predict,err :%v", err)
		h.Response.Status = 1
		h.Response.Message = err.Error()
		return
	}
	h.Response.Data.Order = torchPredictResult.Order
}

func (h *PredictHandler) CheckParam() error {

	return nil
}

func ConvertInputFromRpcToTorch(input *post_model_predict.PredictData) *model.ModelPredictInput {
	var features [][]float32
	for _, point := range input.Points {
		features = append(features, point.Features)
	}
	torchInput := &model.ModelPredictInput{
		Feature: features,
		Start:   input.Start,
	}
	return torchInput
}
