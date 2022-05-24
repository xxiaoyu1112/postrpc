package main

import (
	"log"
	"net"

	post_model_predict "post_model_predict/idl/post_model_predict"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
	port = ":50054"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	post_model_predict.RegisterPostModelPredictServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
