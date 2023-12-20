package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	gencontent "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/content"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type ICollectionService interface {
	GetCatPreviews(ctx context.Context, req *core_api.GetCatPreviewsReq) (*core_api.GetCatPreviewsResp, error)
	GetCatDetail(ctx context.Context, req *core_api.GetCatDetailReq) (*core_api.GetCatDetailResp, error)
	NewCat(ctx context.Context, req *core_api.NewCatReq) (*core_api.NewCatResp, error)
	DeleteCat(ctx context.Context, req *core_api.DeleteCatReq) (*core_api.DeleteCatResp, error)
	CreateImage(ctx context.Context, req *core_api.CreateImageReq) (*core_api.CreateImageResp, error)
	DeleteImage(ctx context.Context, req *core_api.DeleteImageReq) (*core_api.DeleteImageResp, error)
	GetImageByCat(ctx context.Context, req *core_api.GetImageByCatReq) (*core_api.GetImageByCatResp, error)
}

type CollectionService struct {
	MeowchatContent       meowchat_content.IMeowchatContent
	Config                *config.Config
	PlatformSts           platform_sts.IPlatformSts
	CatImageDomainService service.ICatImageDomainService
}

var CollectionServiceSet = wire.NewSet(
	wire.Struct(new(CollectionService), "*"),
	wire.Bind(new(ICollectionService), new(*CollectionService)),
)

func (s *CollectionService) GetCatPreviews(ctx context.Context, req *core_api.GetCatPreviewsReq) (*core_api.GetCatPreviewsResp, error) {
	resp := new(core_api.GetCatPreviewsResp)
	pageSize := consts.DefaultPageSize
	if req.PaginationOption == nil {
		req.PaginationOption = &basic.PaginationOptions{}
	}
	if req.GetKeyword() == "" {
		data, err := s.MeowchatContent.ListCat(ctx, &gencontent.ListCatReq{
			CommunityId: req.CommunityId,
			Count:       pageSize,
			Skip:        req.PaginationOption.GetPage() * pageSize,
		})
		if err != nil {
			return nil, err
		}
		resp.Total = data.Total
		resp.Cats = make([]*core_api.CatPreview, 0, pageSize)
		err = copier.Copy(&resp.Cats, data.Cats)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(resp.Cats); i++ {
			if len(data.Cats[i].Avatars) > 0 {
				resp.Cats[i].AvatarUrl = data.Cats[i].Avatars[0]
			}
		}
	} else {
		data, err := s.MeowchatContent.SearchCat(ctx, &gencontent.SearchCatReq{
			CommunityId: req.CommunityId,
			Count:       pageSize,
			Skip:        *req.PaginationOption.Page * pageSize,
			Keyword:     req.GetKeyword(),
		})
		if err != nil {
			return nil, err
		}
		resp.Total = data.Total
		resp.Cats = make([]*core_api.CatPreview, 0, pageSize)
		err = copier.Copy(&resp.Cats, data.Cats)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(resp.Cats); i++ {
			if len(data.Cats[i].Avatars) > 0 {
				resp.Cats[i].AvatarUrl = data.Cats[i].Avatars[0]
			}
		}
	}

	return resp, nil
}

func (s *CollectionService) GetCatDetail(ctx context.Context, req *core_api.GetCatDetailReq) (*core_api.GetCatDetailResp, error) {
	resp := new(core_api.GetCatDetailResp)
	data, err := s.MeowchatContent.RetrieveCat(ctx, &gencontent.RetrieveCatReq{
		CatId: req.CatId,
	})
	if err != nil {
		return nil, err
	}

	resp.Cat = new(content.Cat)
	err = copier.Copy(resp.Cat, data.Cat)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CollectionService) NewCat(ctx context.Context, req *core_api.NewCatReq) (*core_api.NewCatResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewCatResp)
	cat := new(gencontent.Cat)

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
		var data *gencontent.CreateCatResp
		data, err = s.MeowchatContent.CreateCat(ctx, &gencontent.CreateCatReq{Cat: cat})
		if err != nil {
			return nil, err
		}
		resp.CatId = data.CatId
	} else {
		_, err = s.MeowchatContent.UpdateCat(ctx, &gencontent.UpdateCatReq{Cat: cat})
		if err != nil {
			return nil, err
		}
		resp.CatId = cat.Id
	}

	return resp, nil
}

func (s *CollectionService) DeleteCat(ctx context.Context, req *core_api.DeleteCatReq) (*core_api.DeleteCatResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteCatResp)
	_, err := s.MeowchatContent.DeleteCat(ctx, &gencontent.DeleteCatReq{CatId: req.CatId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CollectionService) CreateImage(ctx context.Context, req *core_api.CreateImageReq) (*core_api.CreateImageResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.CreateImageResp)

	for i := 0; i < len(req.Images); i++ {
		u, err := url.Parse(req.Images[i].Url)
		if err != nil {
			return nil, err
		}
		u.Host = s.Config.CdnHost
		req.Images[i].Url = u.String()
	}
	i := make([]string, len(req.Images))
	for key, image := range req.Images {
		i[key] = image.Url
	}
	r, err := s.PlatformSts.PhotoCheck(ctx, &sts.PhotoCheckReq{
		User: user,
		Url:  i,
	})
	if err != nil {
		return nil, err
	}
	if r.Pass == false {
		return nil, errors.NewBizError(10002, "PhotoCheck don't pass")
	}

	rpcReq := new(gencontent.CreateImageReq)
	err = copier.Copy(rpcReq, req)
	if err != nil {
		return nil, err
	}

	res, err := s.MeowchatContent.CreateImage(ctx, rpcReq)
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
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteImageResp)
	data := gencontent.DeleteImageReq{ImageId: req.ImageId}
	_, err := s.MeowchatContent.DeleteImage(ctx, &data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *CollectionService) GetImageByCat(ctx context.Context, req *core_api.GetImageByCatReq) (*core_api.GetImageByCatResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	resp := new(core_api.GetImageByCatResp)
	data := gencontent.ListImageReq{
		CatId:    req.CatId,
		Limit:    req.Limit,
		Backward: req.Backward,
	}
	if req.GetPrevId() != "" {
		data.PrevId = req.PrevId
	}
	res, err := s.MeowchatContent.ListImage(ctx, &data)
	if err != nil {
		return nil, err
	}

	resp.Total = res.Total
	resp.Images = make([]*core_api.Image, len(res.GetImages()))
	util.ParallelRun(lo.Map(res.GetImages(), func(image *gencontent.Image, i int) func() {
		return func() {
			img := &core_api.Image{
				Id:    image.Id,
				Url:   image.Url,
				CatId: image.CatId,
			}
			util.ParallelRun(
				func() {
					if user.GetUserId() == "" {
						return
					}
					_ = s.CatImageDomainService.LoadIsCurrentUserLiked(ctx, img, user.GetUserId())
				},
				func() {
					_ = s.CatImageDomainService.LoadLikeCount(ctx, img)
				})
			resp.Images[i] = img
		}
	})...)
	return resp, nil
}
