package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"post_data_manage/model"
)

func GetPostPredictDeal(ctx context.Context, date, region string, pageSize, pageNum int64) ([]*model.PostPredict, error) {
	var findoptions *options.FindOptions = &options.FindOptions{}
	if pageSize > 0 {
		findoptions.SetLimit(pageSize)
		findoptions.SetSkip(pageSize * pageNum)
	}
	cur, err := postPredictDeal.Find(ctx, bson.M{"deal_date": date, "region": region}, findoptions)
	if err != nil {
		log.Printf("[GetPostPredictDeal] error call mongo find,err: %v", err)
		return nil, err
	}
	var deals []*model.PostPredict
	err = cur.All(ctx, &deals)
	if err != nil {
		log.Printf("[GetPostPredictDeal] error call mongo curosr,err: %v", err)
		return nil, err
	}
	return deals, nil
}
