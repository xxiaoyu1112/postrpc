package main

import (
	"log"
	"net"
	"post_data_manage/idl/post_data_manage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	post_data_manage.RegisterPostDataManageServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// ctx := context.Background()
	// res, _ := mongo.GetPostPredictDeal(ctx, "20200423", "上海市", 10, 10)
	// log.Printf("%+v", res)
}
