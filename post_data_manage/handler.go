package main

import (
	"context"
	Handler "post_data_manage/handler"
	"post_data_manage/idl/post_data_manage"
)

type Server struct {
	post_data_manage.UnsafePostDataManageServer
}

func (server *Server) GeneratePostPredictData(ctx context.Context, req *post_data_manage.GeneratePostPredictDataRequest) (*post_data_manage.GeneratePostPredictDataResponse, error) {
	handler := Handler.NewGeneratePostPredictDataHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) GetPostPredictData(ctx context.Context, req *post_data_manage.GetPostPredictDataRequest) (*post_data_manage.GetPostPredictDataResponse, error) {
	handler := Handler.NewGetPostPredictDataHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) GetPostPredictTaskStatus(ctx context.Context, req *post_data_manage.GetPostPredictTaskStatusRequest) (*post_data_manage.GetPostPredictTaskStatusResponse, error) {
	handler := Handler.NewGetPostPredictTaskStatusHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) GetRawDataTree(ctx context.Context, req *post_data_manage.GetRawDataTreeRequest) (*post_data_manage.GetRawDataTreeResponse, error) {
	handler := Handler.NewGetRawDataTreeHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) GetRawData(ctx context.Context, req *post_data_manage.GetRawDataRequest) (*post_data_manage.GetRawDataResponse, error) {
	handler := Handler.NewGetRawDataHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}
