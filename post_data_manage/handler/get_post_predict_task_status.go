package handler

import (
	"context"
	"log"
	"post_data_manage/idl/post_data_manage"
	"post_data_manage/model"
	"post_data_manage/mysql"
)

type GetPostPredictTaskStatusHandler struct {
	Request  *post_data_manage.GetPostPredictTaskStatusRequest
	Response *post_data_manage.GetPostPredictTaskStatusResponse
	Ctx      context.Context
}

func NewGetPostPredictTaskStatusHandler(Ctx context.Context, Request *post_data_manage.GetPostPredictTaskStatusRequest) *GetPostPredictTaskStatusHandler {
	return &GetPostPredictTaskStatusHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_data_manage.GetPostPredictTaskStatusResponse{
			Message: "success",
			Status:  0,
			PredictTaskStatus: &post_data_manage.PredictTaskStatus{
				Status:   0, // not start
				Finished: false,
			},
		},
	}
}
func (h *GetPostPredictTaskStatusHandler) Run() {
	task, err := h.GetTask()
	if err != nil {
		log.Printf("[GetPostPredictTaskStatusHandler] get task err : %+v", err)
		h.Response.Status = 1
		h.Response.Message = "error"
		return
	}
	if task == nil {
		h.Response.PredictTaskStatus.Status = 0
	} else {
		h.Response.PredictTaskStatus.Status = task.PostTaskStatus
	}
}

func (h *GetPostPredictTaskStatusHandler) CheckParam() error {
	return nil
}

func (h *GetPostPredictTaskStatusHandler) GetTask() (*model.PostTask, error) {
	taskId := h.Request.GetTaskId()
	task, err := mysql.GetTaskById(h.Ctx, taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}
