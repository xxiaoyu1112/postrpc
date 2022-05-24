package api_test

import (
	"log"
	"post_model_manage/api"
	"testing"
)

func TestGetModelStates(t *testing.T) {
	_, err := api.GetModelStates("greedy_distance")
	if err != nil {
		log.Printf("err:%v \n",err)
		return
	}else{
		log.Print("pass ")
	}
}
