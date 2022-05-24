package mongo

import (
	"context"
	"log"
	"post_data_manage/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRawDataByTag(ctx context.Context, tag string, pageSize, pageNum int64) ([]string, error) {
	var findoptions *options.FindOptions = &options.FindOptions{}
	if pageSize > 0 {
		findoptions.SetLimit(pageSize)
		findoptions.SetSkip(pageSize * pageNum)
	}
	cur, err := postRawDeal.Find(ctx, bson.M{"tag": tag}, findoptions)
	if err != nil {
		log.Printf("[GetRawDataByTag] error call mongo find,err: %v", err)
		return nil, err
	}
	var deals []*model.PostDeal
	err = cur.All(ctx, &deals)
	if err != nil {
		log.Printf("[GetRawDataByTag] error call mongo curosr,err: %v", err)
		return nil, err
	}
	return nil, nil
}
