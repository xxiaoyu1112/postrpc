package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"post_model_manage/idl/post_model_manage"
	"post_model_manage/model"
	pb "post_model_manage/rpc/rpc_idl/torch_serve"

	"google.golang.org/grpc"
)

const (
	management_address = "211.71.76.189:7071"
)

var management_client pb.ManagementAPIsServiceClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(management_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	management_client = pb.NewManagementAPIsServiceClient(conn)
}

func RegisterModel(ctx context.Context, modelUrl, modelName string) (string, error) {
	req := &pb.RegisterModelRequest{
		InitialWorkers: 1,
		ModelName:      modelName,
		Synchronous:    false, //是否是同步加载,我们选择异步加载
		Url:            modelUrl,
	}
	resp, err := management_client.RegisterModel(ctx, req)
	if err != nil {
		log.Printf("[Error] Error call rpc RegisterModel, err: %v", err)
		return err.Error(), err
	}
	return resp.Msg, nil
}

func GetModelList(ctx context.Context) (*model.ModelListResponse, error) {
	req := &pb.ListModelsRequest{
		Limit:         100,
		NextPageToken: 0,
	}
	resp, err := management_client.ListModels(ctx, req)
	if err != nil {
		log.Printf("[Error] Error call rpc ListModels, err: %v", err)
		return nil, err
	}
	var modelList model.ModelListResponse
	if err = json.Unmarshal([]byte(resp.Msg), &modelList); err != nil {
		log.Printf("[Error] Error call marshal err: %v", err)
		return nil, err
	}
	return &modelList, nil
}

func RemoveModel(ctx context.Context, modelName, modelVersion string) error {
	req := &pb.UnregisterModelRequest{
		ModelName:    modelName,
		ModelVersion: modelVersion,
	}
	_, err := management_client.UnregisterModel(ctx, req)
	if err != nil {
		log.Printf("[Error] Error call rpc UnregisterModel,err:%v", err)
		return err
	}
	return nil
}

func UpdateModelConfig(ctx context.Context, modelName, modelVersion string, maxWorker, minWorker int32) error {
	req := &pb.ScaleWorkerRequest{
		MaxWorker:    maxWorker,
		MinWorker:    minWorker,
		ModelName:    modelName,
		ModelVersion: modelVersion,
	}

	_, err := management_client.ScaleWorker(ctx, req)
	if err != nil {
		log.Printf("[Error] Error call rpc ScaleWorker,err:%v", err)
		return err
	}
	return nil
}

func GetModelStateById(ctx context.Context, id *post_model_manage.ModelIdentity) ([]*model.ModelState, error) {
	req := &pb.DescribeModelRequest{
		ModelName: id.ModelName,
	}
	describeModel, err := management_client.DescribeModel(ctx, req)
	if err != nil {
		log.Printf("[Error] error call rpc GetModelState, err:%v", err)
		return nil, err
	}
	msg := describeModel.Msg
	var modelStates []*model.ModelState
	err = json.Unmarshal([]byte(msg), &modelStates)
	if err != nil {
		log.Printf("[Error] Error call marshal err: %v", err)
		return nil, err
	}
	return modelStates, nil
}

// GetModelStates need to opt in feature
func GetModelStatesByIds(ctx context.Context, identities []*post_model_manage.ModelIdentity) ([]*model.ModelState, error) {
	var modelStates []*model.ModelState
	for _, id := range identities {
		// 3 times for retry if err occur
		for i := 0; i < 3; i++ {
			currentModelStates, err := GetModelStateById(ctx, id)
			if err == nil {
				modelStates = append(modelStates, currentModelStates...)
				break
			} else if i == 2 {
				return nil, fmt.Errorf("Error in GetModelStates, id: %v, err: %v ", id, err)
			}
		}
	}
	return modelStates, nil
}
