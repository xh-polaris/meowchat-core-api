package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/collection"

	collection2 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/collection"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_collection"
)

type ICollectionService interface {
	GetCatPreviews(ctx context.Context, req *core_api.GetCatPreviewsReq) (*core_api.GetCatPreviewsResp, error)
	GetCatDetail(ctx context.Context, req *core_api.GetCatDetailReq) (*core_api.GetCatDetailResp, error)
	NewCat(ctx context.Context, req *core_api.NewCatReq) (*core_api.NewCatResp, error)
	SearchCat(ctx context.Context, req *core_api.SearchCatReq) (*core_api.SearchCatResp, error)
	DeleteCat(ctx context.Context, req *core_api.DeleteCatReq) (*core_api.DeleteCatResp, error)
	CreateImage(ctx context.Context, req *core_api.CreateImageReq) (*core_api.CreateImageResp, error)
	DeleteImage(ctx context.Context, req *core_api.DeleteImageReq) (*core_api.DeleteImageResp, error)
	GetImageByCat(ctx context.Context, req *core_api.GetImageByCatReq) (*core_api.GetImageByCatResp, error)
}

type CollectionService struct {
	Collection meowchat_collection.IMeowchatCollection
	Config     *config.Config
}

func (s *CollectionService) GetCatPreviews(ctx context.Context, req *core_api.GetCatPreviewsReq) (*core_api.GetCatPreviewsResp, error) {
	resp := new(core_api.GetCatPreviewsResp)
	pageSize := consts.DefaultPageSize
	data, err := s.Collection.ListCat(ctx, &collection.ListCatReq{
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

func (s *CollectionService) GetCatDetail(ctx context.Context, req *core_api.GetCatDetailReq) (*core_api.GetCatDetailResp, error) {
	resp := new(core_api.GetCatDetailResp)
	data, err := s.Collection.RetrieveCat(ctx, &collection.RetrieveCatReq{
		CatId: req.CatId,
	})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Cat, data.Cat)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CollectionService) NewCat(ctx context.Context, req *core_api.NewCatReq) (*core_api.NewCatResp, error) {
	resp := new(core_api.NewCatResp)
	cat := new(collection.Cat)

	for i := 0; i < len(req.Avatars); i++ {
		u, err := url.Parse(req.Avatars[i])
		if err != nil {
			return nil, err
		}
		u.Host = s.Config.CdnHost
		req.Avatars[i] = u.String()
	}

	err := copier.Copy(cat, req)
	if err != nil {
		return nil, err
	}

	if req.GetId() == "" {
		var data *collection.CreateCatResp
		data, err = s.Collection.CreateCat(ctx, &collection.CreateCatReq{Cat: cat})
		if err != nil {
			return nil, err
		}
		resp.CatId = data.CatId
	} else {
		_, err = s.Collection.UpdateCat(ctx, &collection.UpdateCatReq{Cat: cat})
		if err != nil {
			return nil, err
		}
		resp.CatId = cat.Id
	}

	return resp, nil
}

func (s *CollectionService) SearchCat(ctx context.Context, req *core_api.SearchCatReq) (*core_api.SearchCatResp, error) {
	resp := new(core_api.SearchCatResp)
	data, err := s.Collection.SearchCat(ctx, &collection.SearchCatReq{
		CommunityId: req.CommunityId,
		Count:       consts.DefaultPageSize,
		Skip:        req.GetPaginationOption().GetPage() * consts.DefaultPageSize,
		Keyword:     req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Cats = make([]*core_api.CatPreview, 0, consts.DefaultPageSize)
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

func (s *CollectionService) DeleteCat(ctx context.Context, req *core_api.DeleteCatReq) (*core_api.DeleteCatResp, error) {
	resp := new(core_api.DeleteCatResp)
	_, err := s.Collection.DeleteCat(ctx, &collection.DeleteCatReq{CatId: req.CatId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CollectionService) CreateImage(ctx context.Context, req *core_api.CreateImageReq) (*core_api.CreateImageResp, error) {
	resp := new(core_api.CreateImageResp)

	for i := 0; i < len(req.Images); i++ {
		u, err := url.Parse(req.Images[i].Url)
		if err != nil {
			return nil, err
		}
		u.Host = s.Config.CdnHost
		req.Images[i].Url = u.String()
	}

	rpcReq := new(collection.CreateImageReq)
	err := copier.Copy(rpcReq, req)
	if err != nil {
		return nil, err
	}

	res, err := s.Collection.CreateImage(ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	// 规避错误
	if len(res.ImageIds) > 0 {
		resp.ImageIds = res.ImageIds
	}
	return resp, nil
}

func (s *CollectionService) DeleteImage(ctx context.Context, req *core_api.DeleteImageReq) (*core_api.DeleteImageResp, error) {
	resp := new(core_api.DeleteImageResp)
	data := collection.DeleteImageReq{ImageId: req.ImageId}
	_, err := s.Collection.DeleteImage(ctx, &data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *CollectionService) GetImageByCat(ctx context.Context, req *core_api.GetImageByCatReq) (*core_api.GetImageByCatResp, error) {
	resp := new(core_api.GetImageByCatResp)
	data := collection.ListImageReq{
		CatId:    req.CatId,
		Limit:    req.Limit,
		Backward: req.Backward,
	}
	if req.GetPrevId() != "" {
		data.PrevId = req.PrevId
	}
	res, err := s.Collection.ListImage(ctx, &data)
	if err != nil {
		return nil, err
	}

	resp.Total = res.Total
	resp.Images = make([]*collection2.Image, len(res.Images))
	for i, image := range res.Images {
		resp.Images[i] = &collection2.Image{
			Id:    image.Id,
			Url:   image.Url,
			CatId: image.CatId,
		}
	}
	return resp, nil
}

var CollectionServiceSet = wire.NewSet(
	wire.Struct(new(CollectionService), "*"),
	wire.Bind(new(ICollectionService), new(*CollectionService)),
)
