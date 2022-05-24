package main

import (
	"log"
	"net"

	"post_model_manage/consts"
	post_model_manage "post_model_manage/idl/post_model_manage"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
	port = "0.0.0.0:50053"
)

func main() {
	consts.Init()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	post_model_manage.RegisterPostModelManageServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
