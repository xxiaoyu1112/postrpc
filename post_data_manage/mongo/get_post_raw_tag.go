package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPostDealTag(ctx context.Context) ([]string, error) {
	results, err := postRawDeal.Distinct(context.TODO(), "tag", &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	var tags []string
	// tags = append(tags, results...)
	for _, result := range results {
		tags = append(tags, result.(string))
	}
	return tags, nil
}
