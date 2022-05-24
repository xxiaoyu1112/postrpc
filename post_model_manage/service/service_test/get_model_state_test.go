package service_test

import (
	"context"
	"log"
	"post_model_manage/service"
	"testing"
)

func TestGetModelStates(t *testing.T) {
	ctx := context.Background()
	resp, err := service.GetModelStates(ctx)
	if err != nil {
		log.Printf("err:%v \n", err)
		return
	} else {
		log.Printf("pass:%v ", resp)
	}
}
