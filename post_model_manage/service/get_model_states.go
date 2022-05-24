package service

import (
	"context"
	"fmt"
	"log"
	"post_model_manage/api"
	"post_model_manage/model"
	"post_model_manage/rpc"
)

func GetModelStates(ctx context.Context) ([]*model.ModelState, error) {
	modelList, err := rpc.GetModelList(ctx)
	if err != nil {
		log.Printf("[Error] error call rpc GetModelList:%v", err)
		return nil, err
	}
	var modelNames []string
	for _, modelItem := range modelList.Models {
		modelNames = append(modelNames, modelItem.ModelName)
	}
	// GetModelStates need to opt in feature
	var modelStates []*model.ModelState
	for _, modelName := range modelNames {
		// 3 times for retry if err occur
		for i := 0; i < 3; i++ {
			currentModelStates, err := api.GetModelStates(modelName)
			if err == nil {
				modelStates = append(modelStates, currentModelStates...)
				break
			} else if i == 2 {
				return nil, fmt.Errorf("Error in GetModelStates, modelName: %v, err: %v ", modelName, err)
			}
		}
	}
	return modelStates, nil
}
