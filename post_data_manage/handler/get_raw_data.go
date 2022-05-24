package handler

import (
	"context"
	"log"
	"post_data_manage/idl/post_data_manage"
	"post_data_manage/model"
	"post_data_manage/mongo"
	"strings"
)

type GetRawDataHandler struct {
	Request  *post_data_manage.GetRawDataRequest
	Response *post_data_manage.GetRawDataResponse
	Ctx      context.Context
}

func NewGetRawDataHandler(Ctx context.Context, Request *post_data_manage.GetRawDataRequest) *GetRawDataHandler {
	return &GetRawDataHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_data_manage.GetRawDataResponse{
			Message: "success",
			Status:  0,
			Data: &post_data_manage.RawDataPage{
				Total:    0,
				RawDatas: []*post_data_manage.RawData{},
			},
		},
	}
}
func (h *GetRawDataHandler) Run() {
	rawData, err := h.GetRawData()
	if err != nil {
		log.Printf("[GetPostPredictTaskStatusHandler] get task err : %+v", err)
		h.Response.Status = 1
		h.Response.Message = "error"
		return
	}
	h.Response.Data.RawDatas = h.BuildRpcData(rawData)
}

func (h *GetRawDataHandler) CheckParam() error {
	return nil
}
func (h *GetRawDataHandler) GetRawData() ([]*model.PostDeal, error) {
	return mongo.GetPostRawDeal(h.Ctx, h.Request.Id, int64(h.Request.PageSize), int64(h.Request.PageNum))
}
func (h *GetRawDataHandler) BuildRpcData(rawDatas []*model.PostDeal) []*post_data_manage.RawData {
	var rpcRawDatas []*post_data_manage.RawData
	for _, rawData := range rawDatas {
		res := strings.Split(rawData.Raw, ",")
		rpcRawDatas = append(rpcRawDatas, &post_data_manage.RawData{
			DealDate:                  res[0],
			RegionId:                  res[1],
			City:                      res[2],
			PostManId:                 res[3],
			GetDealTime:               res[4],
			AppointmentTimeOne:        res[5],
			AppointmentTimeTwo:        res[6],
			DealLongitude:             res[7],
			DealLatitude:              res[8],
			DealRegionId:              res[9],
			DealBlockTypeId:           res[10],
			DealBlockType:             res[11],
			DealFinishTime:            res[12],
			RecentFinishDealTime:      res[13],
			RecentFinishDealLongitude: res[14],
			RecentFinishDealLatitude:  res[15],
			FinishDealPrecision:       res[16],
			RecentGetDealTime:         res[17],
			RecentGetDealLongitude:    res[18],
			RecentGetDealLatitude:     res[19],
			GetDealPrecision:          res[20],
		})
	}
	return rpcRawDatas
}
