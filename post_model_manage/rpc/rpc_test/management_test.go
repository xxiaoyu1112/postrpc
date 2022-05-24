package rpc_test

import (
	"context"
	"log"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/rpc"
	"testing"
)

// Unit Test
func TestUpdateModelConfig(t *testing.T) {
	ctx := context.Background()
	err := rpc.UpdateModelConfig(ctx, "greedy_distance", "1.0", 1, 1)
	if err != nil {
		log.Printf("err:%v", err)
	} else {
		log.Print("pass")
	}
}

func TestGetModelState(t *testing.T) {
	ctx := context.Background()
	_, err := rpc.GetModelStateById(ctx, &post_model_manage.ModelIdentity{ModelName: "greedy_distance", ModelVersion: "1.0"})
	if err != nil {
		log.Printf("err:%v", err)
	} else {
		log.Print("pass")
	}
}
