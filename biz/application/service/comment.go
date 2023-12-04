package service

import (
	"context"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/platform/comment"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type ICommentService interface {
	GetComments(ctx context.Context, req *core_api.GetCommentsReq, user *basic.UserMeta) (*core_api.GetCommentsResp, error)
	NewComment(ctx context.Context, req *core_api.NewCommentReq, user *basic.UserMeta) (*core_api.NewCommentResp, error)
	DeleteComment(ctx context.Context, req *core_api.DeleteCommentReq) (*core_api.DeleteCommentResp, error)
}

type CommentService struct {
	Config               *config.Config
	CommentDomainService service.ICommentDomainService
	PlatformComment      platform_comment.IPlatformCommment
	PlatformSts          platform_sts.IPlatformSts
	MeowchatContent      meowchat_content.IMeowchatContent
}

var CommentServiceSet = wire.NewSet(
	wire.Struct(new(CommentService), "*"),
	wire.Bind(new(ICommentService), new(*CommentService)),
)

func (s *CommentService) GetComments(ctx context.Context, req *core_api.GetCommentsReq, userMeta *basic.UserMeta) (*core_api.GetCommentsResp, error) {
	resp := new(core_api.GetCommentsResp)

	pageSize := int64(10)
	data, err := s.PlatformComment.ListCommentByParent(ctx, &gencomment.ListCommentByParentReq{
		Id:             req.Id,
		Type:           gencomment.CommentType(req.Type),
		Skip:           req.Page * pageSize,
		Limit:          pageSize,
		OnlyFirstLevel: lo.ToPtr(true),
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Comments = make([]*core_api.Comment, len(data.Comments))
	// 并发获取额外信息
	util.ParallelRun(lo.Map(data.Comments, func(item *gencomment.Comment, i int) func() {
		return func() {
			c := &core_api.Comment{
				Id:       item.Id,
				CreateAt: item.CreateAt,
				Text:     item.Text,
			}
			util.ParallelRun([]func(){
				func() {
					_ = s.CommentDomainService.LoadIsCurrentUserLiked(ctx, c, userMeta.UserId)
				},
				func() {
					if item.ReplyTo == "" {
						return
					}
					_ = s.CommentDomainService.LoadReplyUser(ctx, c, item.ReplyTo)
				},
				func() {
					_ = s.CommentDomainService.LoadCommentCount(ctx, c)
				},
				func() {
					_ = s.CommentDomainService.LoadLikeCount(ctx, c)
				},
				func() {
					_ = s.CommentDomainService.LoadAuthor(ctx, c, item.AuthorId)
				},
			})
			resp.Comments[i] = c
		}
	}))

	return resp, nil
}

func (s *CommentService) NewComment(ctx context.Context, req *core_api.NewCommentReq, user *basic.UserMeta) (*core_api.NewCommentResp, error) {
	resp := new(core_api.NewCommentResp)
	r, err := s.PlatformSts.TextCheck(ctx, &sts.TextCheckReq{
		Text:  req.Text,
		User:  user,
		Scene: 2,
		Title: &req.Text,
	})
	if err != nil {
		return nil, err
	}
	if r.Pass == false {
		return nil, errors.NewBizError(10001, "TextCheck don't pass")
	}

	//获取回复用户id
	replyToId := ""
	if req.Type == comment.CommentType_CommentType_Comment {
		replyTo, err := s.PlatformComment.RetrieveCommentById(ctx, &gencomment.RetrieveCommentByIdReq{Id: *req.Id})
		if err != nil {
			return nil, err
		}
		replyToId = replyTo.Comment.AuthorId
	}

	data, err := s.PlatformComment.CreateComment(ctx, &gencomment.CreateCommentReq{
		Text:         req.Text,
		FirstLevelId: req.GetFirstLevelId(),
		AuthorId:     user.UserId,
		ReplyTo:      replyToId,
		Type:         gencomment.CommentType(req.Type),
		ParentId:     *req.Id,
	})
	if err != nil {
		return nil, err
	}
	if data.GetGetFish() == true {
		_, err = s.MeowchatContent.AddUserFish(ctx, &content.AddUserFishReq{
			UserId: user.UserId,
			Fish:   s.Config.Fish.Comment[data.GetFishTimes-1],
		})
		if err == nil {
			resp.GetFishNum = s.Config.Fish.Comment[data.GetFishTimes-1]
			resp.GetFishTimes = data.GetGetFishTimes()
		}
	}
	resp.GetFish = data.GetGetFish()
	return resp, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, req *core_api.DeleteCommentReq) (*core_api.DeleteCommentResp, error) {
	resp := new(core_api.DeleteCommentResp)
	_, err := s.PlatformComment.DeleteComment(ctx, &gencomment.DeleteCommentByIdReq{Id: req.CommentId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
