package handler

import (
	"context"
	"log"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/model"
	"post_model_manage/service"
)

type GetModelStatesHandler struct {
	Request  *post_model_manage.GetModelStatesRequest
	Response *post_model_manage.GetModelStatesResponse
	Ctx      context.Context
}

func NewGetModelStatesHandler(Ctx context.Context, Request *post_model_manage.GetModelStatesRequest) *GetModelStatesHandler {
	return &GetModelStatesHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_model_manage.GetModelStatesResponse{
			Message: "success",
			Status:  0,
			Data:    &post_model_manage.GetModelStatesData{},
		},
	}
}
func (h *GetModelStatesHandler) Run() {
	states, err := service.GetModelStates(h.Ctx)
	if err != nil {
		log.Printf("[Error] call service.GetModelStates,err:%v", err)
		return
	}
	h.Response.Data.ModelStates = h.ConvertModelStatesFromTorchToRpc(states)
}

func (h *GetModelStatesHandler) CheckParam() error {

	return nil
}
func (h *GetModelStatesHandler) ConvertModelStatesFromTorchToRpc(torchModelStates []*model.ModelState) []*post_model_manage.ModelState {
	var rpcModelStates []*post_model_manage.ModelState
	for _, torchModelState := range torchModelStates {
		rpcModelState := h.ConvertModelStateFromTorchToRpc(torchModelState)
		rpcModelStates = append(rpcModelStates, rpcModelState)
	}
	return rpcModelStates
}
func (h *GetModelStatesHandler) ConvertModelStateFromTorchToRpc(torchModelState *model.ModelState) *post_model_manage.ModelState {
	var rpcWorkers []*post_model_manage.WORKER
	for _, torchWorker := range torchModelState.Workers {
		rpcWorkers = append(rpcWorkers, &post_model_manage.WORKER{
			Id:          torchWorker.Id,
			StartTime:   torchWorker.StartTime,
			Status:      torchWorker.Status,
			MemoryUsage: int32(torchWorker.MemoryUsage),
			Pid:         int32(torchWorker.Pid),
			Gpu:         torchWorker.Gpu,
			GpuUsage:    torchWorker.GpuUsage,
		})
	}
	rpcModelState := &post_model_manage.ModelState{
		ModelName:       torchModelState.ModelName,
		ModelVersion:    torchModelState.ModelVersion,
		ModelUrl:        torchModelState.ModelUrl,
		Runtime:         torchModelState.Runtime,
		MinWorkers:      int32(torchModelState.MinWorkers),
		MaxWorkers:      int32(torchModelState.MaxWorkers),
		BatchSize:       int32(torchModelState.BatchSize),
		MaxBatchDelay:   int32(torchModelState.MaxBatchDelay),
		LoadedAtStartup: torchModelState.LoadedAtStartup,
		Workers:         rpcWorkers,
	}
	return rpcModelState
}
