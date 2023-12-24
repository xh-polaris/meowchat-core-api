package service

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type ICommentService interface {
	GetComments(ctx context.Context, req *core_api.GetCommentsReq) (*core_api.GetCommentsResp, error)
	NewComment(ctx context.Context, req *core_api.NewCommentReq) (*core_api.NewCommentResp, error)
	DeleteComment(ctx context.Context, req *core_api.DeleteCommentReq) (*core_api.DeleteCommentResp, error)
}

type CommentService struct {
	Config               *config.Config
	CommentDomainService service.ICommentDomainService
	PlatformComment      platform_comment.IPlatformComment
	PlatformSts          platform_sts.IPlatformSts
	MeowchatContent      meowchat_content.IMeowchatContent
	MeowchatSystem       meowchat_system.IMeowchatSystem
	MeowchatUser         meowchat_user.IMeowchatUser
}

var CommentServiceSet = wire.NewSet(
	wire.Struct(new(CommentService), "*"),
	wire.Bind(new(ICommentService), new(*CommentService)),
)

func (s *CommentService) GetComments(ctx context.Context, req *core_api.GetCommentsReq) (*core_api.GetCommentsResp, error) {
	resp := new(core_api.GetCommentsResp)

	userMeta := adaptor.ExtractUserMeta(ctx)
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
			util.ParallelRun(
				func() {
					if userMeta.GetUserId() == "" {
						return
					}
					_ = s.CommentDomainService.LoadIsCurrentUserLiked(ctx, c, userMeta.UserId)
				},
				func() {
					if item.GetReplyTo() == "" {
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
				})
			resp.Comments[i] = c
		}
	})...)

	return resp, nil
}

func (s *CommentService) NewComment(ctx context.Context, req *core_api.NewCommentReq) (*core_api.NewCommentResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
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

	data, err := s.PlatformComment.CreateComment(ctx, &gencomment.CreateCommentReq{
		Text:         req.Text,
		FirstLevelId: req.GetFirstLevelId(),
		AuthorId:     user.UserId,
		ReplyTo:      req.GetReplyToUserId(),
		Type:         gencomment.CommentType(req.Type),
		ParentId:     req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	wechatMessage := &sts.SendMessageReq{
		MessageType:    3,
		TargetUserId:   req.GetReplyToUserId(),
		SourceUserName: "一名喵友",
		SourceContent:  "一条你发过的信息",
		CommentText:    req.Text,
		CreateAt:       time.Now().Unix(),
		User:           user,
	}
	message := &system.Notification{
		TargetUserId:    req.GetReplyToUserId(),
		SourceUserId:    user.UserId,
		SourceContentId: req.GetId(),
		TargetType:      0,
		Type:            system.NotificationType_TypeCommented,
		Text:            req.Text,
		IsRead:          false,
	}

	util.ParallelRun(
		func() {
			u, err := s.MeowchatUser.GetUser(ctx, &genuser.GetUserReq{UserId: user.UserId})
			if err != nil {
				log.CtxError(ctx, "[GetUser] fail, err=%v", err)
				return
			}
			wechatMessage.SourceUserName = u.User.Nickname
		},
		func() {
			if req.GetFirstLevelId() != "" {
				message.TargetType = system.NotificationTargetType_TargetTypeComment
				comment, err := s.PlatformComment.RetrieveCommentById(ctx, &gencomment.RetrieveCommentByIdReq{Id: req.GetFirstLevelId()})
				if err != nil {
					log.CtxError(ctx, "[RetrieveCommentById] fail, err=%v", err)
					return
				}
				wechatMessage.SourceContent = comment.Comment.Text
			} else {
				if req.Type == 2 {
					message.TargetType = system.NotificationTargetType_TargetTypePost
					post, err := s.MeowchatContent.RetrievePost(ctx, &content.RetrievePostReq{PostId: req.GetId()})
					if err != nil {
						log.CtxError(ctx, "[RetrievePost] fail, err=%v", err)
						return
					}
					message.TargetUserId = post.Post.UserId
					wechatMessage.TargetUserId = post.Post.UserId
					wechatMessage.SourceContent = post.Post.GetTitle()
				} else if req.Type == 3 {
					moment, err := s.MeowchatContent.RetrieveMoment(ctx, &content.RetrieveMomentReq{MomentId: req.GetId()})
					if err != nil {
						log.CtxError(ctx, "[RetrieveMoment] fail, err=%v", err)
						return
					}
					message.TargetUserId = moment.Moment.UserId
					wechatMessage.TargetUserId = moment.Moment.UserId
					wechatMessage.SourceContent = moment.Moment.GetTitle()
					message.TargetType = system.NotificationTargetType_TargetTypeMoment
				}
			}
		},
	)

	util.ParallelRun(
		func() {
			_, err = s.MeowchatSystem.AddNotification(ctx, &system.AddNotificationReq{Notification: message})
			if err != nil {
				log.CtxError(ctx, "[AddNotification] Add notification failed, err=%v", err)
				return
			}
		},
		func() {
			_, err = s.PlatformSts.SendMessage(ctx, wechatMessage)
			if err != nil {
				log.CtxError(ctx, "[SendMessage] WechatMessage send failed, err=%v", err)
				return
			}
		},
	)

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
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteCommentResp)
	_, err := s.PlatformComment.DeleteComment(ctx, &gencomment.DeleteCommentByIdReq{Id: req.CommentId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
