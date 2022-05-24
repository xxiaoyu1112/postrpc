package handler

import (
	"context"
	"log"
	"post_data_manage/idl/post_data_manage"
	"post_data_manage/model"
	"post_data_manage/mongo"
)

type GetPostPredictDataHandler struct {
	Request  *post_data_manage.GetPostPredictDataRequest
	Response *post_data_manage.GetPostPredictDataResponse
	Ctx      context.Context
}

func NewGetPostPredictDataHandler(Ctx context.Context, Request *post_data_manage.GetPostPredictDataRequest) *GetPostPredictDataHandler {
	return &GetPostPredictDataHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_data_manage.GetPostPredictDataResponse{
			Message:         "success",
			Status:          0,
			PostPredictData: &post_data_manage.PostPredictData{},
		},
	}
}
func (h *GetPostPredictDataHandler) Run() {
	mongoPostPredict, err := h.GetPostPredictData()
	if err != nil {
		log.Printf("[GetPostPredictDataHandler] error call GetPostPredictData, err: %v", err)
		h.Response.Message = "error"
		h.Response.Status = 1
		return
	}
	rpcPredictData := h.ConvertPostPredictFromMongoToRpc(mongoPostPredict)
	h.Response.PostPredictData.PredictDeals = rpcPredictData
}

func (h *GetPostPredictDataHandler) CheckParam() error {
	return nil
}

func (h *GetPostPredictDataHandler) GetPostPredictData() ([]*model.PostPredict, error) {
	return mongo.GetPostPredictDeal(h.Ctx, h.Request.DealDate, h.Request.DealRegion, int64(h.Request.PageSize), int64(h.Request.PageNum))
}
func (h *GetPostPredictDataHandler) ConvertPostPredictFromMongoToRpc(mongoPredictDeals []*model.PostPredict) []*post_data_manage.PredictDeal {
	var rpcPredictDeals []*post_data_manage.PredictDeal
	for _, mongoDeal := range mongoPredictDeals {
		var order []int32
		var points []*post_data_manage.Point
		for i := 0; i < mongoDeal.Length; i++ {
			order = append(order, int32(mongoDeal.Order[i]))
			points = append(points, &post_data_manage.Point{
				Features: mongoDeal.Points[i],
			})
		}
		rpcPredictDeal := &post_data_manage.PredictDeal{
			Points:   points,
			Order:    order,
			Start:    int32(mongoDeal.Start),
			Length:   int32(mongoDeal.Length),
			Region:   mongoDeal.Region,
			DealDate: mongoDeal.DealDate,
		}
		rpcPredictDeals = append(rpcPredictDeals, rpcPredictDeal)
	}
	return rpcPredictDeals
}
