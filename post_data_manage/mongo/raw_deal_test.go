package mongo

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestGetRawDeal(t *testing.T) {
	fmt.Println("start")
	ctx := context.Background()
	res, _ := GetPostRawDeal(ctx, "上海市-20200423", 10, 1)
	r := strings.Split(res[0].Raw, ",")
	fmt.Printf("%+v", len(r))
}
func TestGetRawTag(t *testing.T) {
	fmt.Print("start")
	ctx := context.Background()
	GetPostDealTag(ctx)
}
