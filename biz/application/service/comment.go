package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	user2 "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"
)

const pageSize = 10

type ICommentService interface {
	GetComments(ctx context.Context, req *core_api.GetCommentsReq) (*core_api.GetCommentsResp, error)
	NewComment(ctx context.Context, req *core_api.NewCommentReq, user *basic.UserMeta) (*core_api.NewCommentResp, error)
	DeleteComment(ctx context.Context, req *core_api.DeleteCommentReq) (*core_api.DeleteCommentResp, error)
}

type CommentService struct {
	Config  *config.Config
	Comment platform_comment.IPlatformCommment
	User    meowchat_user.IMeowchatUser
	Sts     platform_sts.IPlatformSts
}

var CommentServiceSet = wire.NewSet(
	wire.Struct(new(CommentService), "*"),
	wire.Bind(new(ICommentService), new(*CommentService)),
)

func (s *CommentService) GetComments(ctx context.Context, req *core_api.GetCommentsReq) (*core_api.GetCommentsResp, error) {
	resp := new(core_api.GetCommentsResp)

	data, err := s.Comment.ListCommentByParent(ctx, &gencomment.ListCommentByParentReq{
		ParentId: req.Id,
		Type:     req.Scope,
		Skip:     req.Page * pageSize,
		Limit:    pageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Comments = make([]*core_api.Comment, len(data.Comments))
	for i, comment := range data.Comments {
		// 评论作者信息
		var author user.UserPreview
		author.Id = comment.AuthorId
		// 暂时不处理error
		user, err := s.User.GetUser(ctx, &user2.GetUserReq{
			UserId: comment.AuthorId,
		})
		if err != nil {
			return nil, err
		}
		if user != nil && err == nil {
			author.Nickname = user.User.Nickname
			author.AvatarUrl = user.User.AvatarUrl
		}

		// 回复对象用户名
		replyName := ""
		if comment.ReplyTo != "" {
			replyToUser, err := s.User.GetUser(ctx, &user2.GetUserReq{
				UserId: comment.ReplyTo,
			})
			if replyToUser != nil && err == nil {
				replyName = replyToUser.User.Nickname
			}
		}

		// 子评论数量
		// TODO count rpc
		count := 0
		children, err := s.Comment.ListCommentByParent(ctx, &gencomment.ListCommentByParentReq{
			Type:     "comment",
			ParentId: comment.Id,
			Skip:     0,
			Limit:    9999,
		})
		if children != nil && err == nil {
			count = len(children.Comments)
		}

		resp.Comments[i] = &core_api.Comment{
			Id:        comment.Id,
			CreateAt:  comment.CreateAt,
			Text:      comment.Text,
			User:      &author,
			Comments:  int64(count),
			ReplyName: replyName,
		}
	}

	return nil, nil
}

func (s *CommentService) NewComment(ctx context.Context, req *core_api.NewCommentReq, user *basic.UserMeta) (*core_api.NewCommentResp, error) {
	resp := new(core_api.NewCommentResp)
	userId := user.UserId
	openId := user.WechatUserMeta.OpenId

	r, err := s.Sts.TextCheck(ctx, &sts.TextCheckReq{
		Text: req.Text,
		User: &basic.UserMeta{
			WechatUserMeta: &basic.WechatUserMeta{
				OpenId: openId,
			},
		},
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
	if req.Scope == "comment" {
		replyTo, err := s.Comment.RetrieveCommentById(ctx, &gencomment.RetrieveCommentByIdReq{Id: *req.Id})
		if err != nil {
			return nil, err
		}
		replyToId = replyTo.Comment.AuthorId
	}

	_, err = s.Comment.CreateComment(ctx, &gencomment.CreateCommentReq{
		Text:     req.Text,
		AuthorId: userId,
		ReplyTo:  replyToId,
		Type:     req.Scope,
		ParentId: *req.Id,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, req *core_api.DeleteCommentReq) (*core_api.DeleteCommentResp, error) {
	resp := new(core_api.DeleteCommentResp)
	_, err := s.Comment.DeleteComment(ctx, &gencomment.DeleteCommentByIdReq{Id: req.CommentId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
