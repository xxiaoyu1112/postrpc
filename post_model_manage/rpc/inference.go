package rpc

import (
	"log"
	pb "post_model_manage/rpc/rpc_idl/torch_serve"

	"google.golang.org/grpc"
)

const (
	inference_address = "211.71.76.189:7070"
)

var inference_client pb.InferenceAPIsServiceClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(inference_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	inference_client = pb.NewInferenceAPIsServiceClient(conn)
}
