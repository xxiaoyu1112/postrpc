package main

import (
	"context"
	"post_model_manage/handler"
	"post_model_manage/idl/post_model_manage"
)

type Server struct {
	post_model_manage.UnsafePostModelManageServer
}

func (s *Server) GetModels(ctx context.Context, request *post_model_manage.GetModelsRequest) (*post_model_manage.GetModelsResponse, error) {
	Handler := handler.NewGetModelsHandler(ctx, request)
	Handler.Run()
	return Handler.Response, nil
}

func (s *Server) RemoveModel(ctx context.Context, request *post_model_manage.RemoveModelRequest) (*post_model_manage.RemoveModelResponse, error) {
	Handler := handler.NewRemoveModelHandler(ctx, request)
	Handler.Run()
	return Handler.Response, nil
}

func (s *Server) UpdateModelConfig(ctx context.Context, request *post_model_manage.UpdateModelConfigRequest) (*post_model_manage.UpdateModelConfigResponse, error) {
	Handler := handler.NewUpdateModelConfigHandler(ctx, request)
	Handler.Run()
	return Handler.Response, nil
}

func (s *Server) GetModelStates(ctx context.Context, request *post_model_manage.GetModelStatesRequest) (*post_model_manage.GetModelStatesResponse, error) {
	Handler := handler.NewGetModelStatesHandler(ctx, request)
	Handler.Run()
	return Handler.Response, nil
}

/*  */
func (s *Server) RegisterModel(ctx context.Context, request *post_model_manage.RegisterModelRequest) (*post_model_manage.RegisterModelResponse, error) {
	Handler := handler.NewRegisterModelHandler(ctx, request)
	Handler.Run()
	return Handler.Response, nil
}
