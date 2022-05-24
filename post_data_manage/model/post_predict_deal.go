package model

type PostPredict struct {
	Points   [][]float64 `bson:"points"`
	Order    []int       `bson:"label"`
	Start    int         `bson:"start"`
	Length   int         `bson:"length"`
	Region   string      `bson:"region"`
	DealDate string      `bson:"deal_date"`
}
