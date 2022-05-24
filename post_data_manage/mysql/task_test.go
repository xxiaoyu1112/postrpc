package mysql

import (
	"context"
	"fmt"
	"post_data_manage/model"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ctx := context.Background()
	task := &model.PostTask{
		PostTaskId:     "uuid",
		PostTaskDate:   time.Now(),
		PostTaskStatus: 0,
	}
	err := CreateTask(ctx, task)
	fmt.Printf("%+v", err)
}

func TestSelect(t *testing.T) {
	ctx := context.Background()
	res, err := GetTaskById(ctx, "uuisd")
	if err != nil {
		fmt.Printf("\n %+v \n", err)
	} else {
		fmt.Printf("\n %+v \n", res)
	}

}
