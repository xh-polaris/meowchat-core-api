package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	pb2 "github.com/xh-polaris/meowchat-user-rpc/pb"
)

const pageSize = 10

type ICommentService interface {
	GetComments(ctx context.Context, req *core_api.GetCommentsReq) (*core_api.GetCommentsResp, error)
	NewComment(ctx context.Context, req *core_api.NewCommentReq) (*core_api.NewCommentResp, error)
	DeleteComment(ctx context.Context, req *core_api.DeleteCommentReq) (*core_api.DeleteCommentResp, error)
}

type CommentService struct {
	Config  *config.Config
	Comment platform_comment.IPlatformCommment
	User    meowchat_user.IMeowchatUser
}

var CommentServiceSet = wire.NewSet(
	wire.Struct(new(CommentService), "*"),
	wire.Bind(new(ICommentService), new(*CommentService)),
)

func (s *CommentService) GetComments(ctx context.Context, req *core_api.GetCommentsReq) (*core_api.GetCommentsResp, error) {
	resp := new(core_api.GetCommentsResp)

	data, err := s.Comment.ListCommentByParent(ctx, &pb.ListCommentByParentRequest{
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
		user, err := s.User.GetUser(ctx, &pb2.GetUserReq{
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
			replyToUser, err := s.User.GetUser(ctx, &pb2.GetUserReq{
				UserId: comment.ReplyTo,
			})
			if replyToUser != nil && err == nil {
				replyName = replyToUser.User.Nickname
			}
		}

		// 子评论数量
		// TODO count rpc
		count := 0
		children, err := s.Comment.ListCommentByParent(ctx, &pb.ListCommentByParentRequest{
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

func (s *CommentService) NewComment(ctx context.Context, req *core_api.NewCommentReq) (*core_api.NewCommentResp, error) {
	resp := new(core_api.NewCommentResp)
	userId := ctx.Value("userId").(string)
	//openId := ctx.Value("openId").(string)
	//
	//err := util.MsgSecCheck(ctx, s, req.Text, openId, 2)
	//if err != nil {
	//	return nil, err
	//}
	// 获取回复用户id
	replyToId := ""
	if req.Scope == "comment" {
		replyTo, err := s.Comment.RetrieveCommentById(ctx, &pb.RetrieveCommentByIdRequest{Id: *req.Id})
		if err != nil {
			return nil, err
		}
		replyToId = replyTo.Comment.AuthorId
	}

	_, err := s.Comment.CreateComment(ctx, &pb.CreateCommentRequest{
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
	_, err := s.Comment.DeleteComment(ctx, &pb.DeleteCommentByIdRequest{Id: req.CommentId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
