package handler

import (
	"context"
	"log"
	"post_data_collect/idl/post_data_collect"
	"post_data_collect/model"
	"post_data_collect/mq"
)

type CollectPostDealHandler struct {
	Request  *post_data_collect.CollectPostDealRequest
	Response *post_data_collect.CollectPostDealResponse
	Ctx      context.Context
}

func NewCollectPostDealHandler(Ctx context.Context, Request *post_data_collect.CollectPostDealRequest) *CollectPostDealHandler {
	return &CollectPostDealHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_data_collect.CollectPostDealResponse{
			Message: "success",
			Status:  0,
		},
	}
}
func (h *CollectPostDealHandler) Run() {
	task := h.BuildCollectTask()
	err := h.SendTask(task)
	if err != nil {
		log.Printf("[CollectPostDealHandler] call  h.SendTask error,err : %v", err)
	}
}

func (h *CollectPostDealHandler) CheckParam() error {

	return nil
}

func (h *CollectPostDealHandler) BuildCollectTask() *model.CollectTask {
	return &model.CollectTask{
		TaskId:      h.Request.PostDealId,
		TaskContent: h.Request.PostDealContent,
	}
}
func (h *CollectPostDealHandler) SendTask(task *model.CollectTask) error {
	return mq.SendCollectTask(h.Ctx, task)
}
