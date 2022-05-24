package handler

import (
	"context"
	"log"
	"post_data_manage/idl/post_data_manage"
	"post_data_manage/mongo"
	"strings"
)

type GetRawDataTreeHandler struct {
	Request  *post_data_manage.GetRawDataTreeRequest
	Response *post_data_manage.GetRawDataTreeResponse
	Ctx      context.Context
}

func NewGetRawDataTreeHandler(Ctx context.Context, Request *post_data_manage.GetRawDataTreeRequest) *GetRawDataTreeHandler {
	return &GetRawDataTreeHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_data_manage.GetRawDataTreeResponse{
			Message: "success",
			Status:  0,
			Tree: &post_data_manage.RawDataTree{
				Root: nil,
			},
		},
	}
}
func (h *GetRawDataTreeHandler) Run() {
	tag, err := h.GetTag()
	if err != nil {
		log.Printf("[GetPostPredictTaskStatusHandler] get task err : %+v", err)
		h.Response.Status = 1
		h.Response.Message = "error"
		return
	}
	h.BuildResp(tag)
}

func (h *GetRawDataTreeHandler) CheckParam() error {
	return nil
}
func (h *GetRawDataTreeHandler) GetTag() (map[string][]string, error) {
	tags, err := mongo.GetPostDealTag(h.Ctx)
	if err != nil {
		return nil, err
	}
	pairs := make(map[string][]string)
	for _, tag := range tags {
		pair := strings.Split(tag, "-")
		dates, contains := pairs[pair[0]]
		if contains {
			dates = append(dates, pair[1])
		} else {
			dates = []string{pair[1]}
		}
		pairs[pair[0]] = dates
	}
	return pairs, nil
}

func (h *GetRawDataTreeHandler) BuildResp(m map[string][]string) {
	var children []*post_data_manage.RawDataTreeNode
	for dealRegion, dealDates := range m {
		var leafNode []*post_data_manage.RawDataTreeNode
		curNode := &post_data_manage.RawDataTreeNode{
			Label:  dealRegion,
			Value:  dealRegion,
			Source: "MongoDB",
			Child:  nil,
		}

		for _, dealDate := range dealDates {
			leafNode = append(leafNode, &post_data_manage.RawDataTreeNode{
				Label:  dealDate,
				Value:  dealRegion + "-" + dealDate,
				Source: "MongoDB",
				Child:  nil,
			})
		}
		curNode.Child = leafNode
		children = append(children, curNode)
	}
	root := &post_data_manage.RawDataTreeNode{
		Label: "MongoDB",
		Value: "MongoDB",
		Child: children,
	}
	h.Response.Tree.Root = root
}
