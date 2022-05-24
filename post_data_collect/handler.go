package main

import (
	"context"
	Handler "post_data_collect/handler"
	"post_data_collect/idl/post_data_collect"
)

type Server struct {
	post_data_collect.UnsafePostDataCollectServer
}

func (server *Server) CollectPostDeal(ctx context.Context, req *post_data_collect.CollectPostDealRequest) (*post_data_collect.CollectPostDealResponse, error) {
	handler := Handler.NewCollectPostDealHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}
