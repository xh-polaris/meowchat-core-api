package service

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_collection"
)

type ICollectionService interface {
	GetCatPreviews(ctx context.Context, req *core_api.GetCatPreviewsReq) (*core_api.GetCatPreviewsResp, error)
}

type CollectionService struct {
	Collection meowchat_collection.IMeowchatCollection
}

var CollectionServiceSet = wire.NewSet(
	wire.Struct(new(CollectionService), "*"),
	wire.Bind(new(ICollectionService), new(*CollectionService)),
)

func (s *CollectionService) GetCatPreviews(ctx context.Context, req *core_api.GetCatPreviewsReq) (*core_api.GetCatPreviewsResp, error) {
	resp := new(core_api.GetCatPreviewsResp)
	pageSize := consts.DefaultPageSize
	data, err := s.Collection.ListCat(ctx, &pb.ListCatReq{
		CommunityId: req.CommunityId,
		Count:       pageSize,
		Skip:        req.Page * pageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Cats = make([]*core_api.CatPreview, 0, pageSize)
	err = copier.Copy(&resp.Cats, data.Cats)
	for i := 0; i < len(resp.Cats); i++ {
		if len(data.Cats[i].Avatars) > 0 {
			resp.Cats[i].AvatarUrl = data.Cats[i].Avatars[0]
		}
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
