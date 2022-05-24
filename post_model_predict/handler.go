package main

import (
	"context"
	Handler "post_model_predict/handler"
	"post_model_predict/idl/post_model_predict"
)

// Server  is used to implement post_model_manage.PostModelManageServer.
type Server struct {
	post_model_predict.UnsafePostModelPredictServer
}

func (server *Server) Predict(ctx context.Context, req *post_model_predict.PredictRequest) (*post_model_predict.PredictResponse, error) {
	handler := Handler.NewPredictHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) PredictAysnc(ctx context.Context, req *post_model_predict.PredictAysncRequest) (*post_model_predict.PredictAysncResponse, error) {
	handler := Handler.NewPredictAysncHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) GetPredictAysncResult(ctx context.Context, req *post_model_predict.GetPredictAysncResultRequest) (*post_model_predict.GetPredictAysncResultResponse, error) {
	handler := Handler.NewGetPredictAysncResultHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}
