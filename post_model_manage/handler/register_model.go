package handler

import (
	"context"
	"fmt"
	"post_model_manage/consts"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/rpc"
	"strings"
)

type RegisterModelHandler struct {
	Request  *post_model_manage.RegisterModelRequest
	Response *post_model_manage.RegisterModelResponse
	Ctx      context.Context
}

func NewRegisterModelHandler(Ctx context.Context, Request *post_model_manage.RegisterModelRequest) *RegisterModelHandler {
	return &RegisterModelHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_manage.RegisterModelResponse{
			Message: "success",
			Status:  0,
		},
	}
}
func (h *RegisterModelHandler) Run() {
	err := h.CheckParam()
	if err != nil {
		h.Response.Status = 1
		h.Response.Message = err.Error()
		return
	}
	modelUrl := h.Request.Url
	modelName := h.getModelName()
	toast, err := rpc.RegisterModel(h.Ctx, modelUrl, modelName)
	if err != nil {
		h.Response.Status = 1
		h.Response.Message = err.Error()
		return
	}
	h.Response.Message = toast
}
func (h *RegisterModelHandler) getModelName() string {
	modelUrl := h.Request.Url
	modelName := ""
	marName := ""
	if consts.HttpMarRegexp.Match([]byte(modelUrl)) {
		httpPath := strings.Split(modelUrl, "/")
		marName = httpPath[len(httpPath)-1]
	} else {
		marName = modelUrl
	}
	// ".mar"
	modelName = marName[:len(marName)-4]
	return modelName
}
func (h *RegisterModelHandler) CheckParam() error {
	modelUrl := h.Request.Url
	if consts.HttpMarRegexp.Match([]byte(modelUrl)) {
		return nil
	}
	if consts.LocalMarRegexp.Match([]byte(modelUrl)) {
		return nil
	}
	return fmt.Errorf("modelUrl is valid:%s", modelUrl)
}
