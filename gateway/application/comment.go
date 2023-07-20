package application

import (
	"context"
	"github.com/pkg/errors"
	"runedance/gateway/rpc"
	"runedance/kitexGen/kitex_gen/commentproto"
	"runedance/kitexGen/kitex_gen/userproto"
	"runedance/models/pojo"
)

var CommentAppIns *CommentAppService

type CommentAppService struct {
}

func NewCommentAppService() *CommentAppService {
	return &CommentAppService{}
}

// CreateComment
// create a comment
func (c CommentAppService) CreateComment(ctx context.Context, appUserID int64, videoID int64, content string) (comment *pojo.Comment, err error) {
	//panic("implement me")
	req := &commentproto.CreateCommentReq{
		UserId:  appUserID,
		VideoId: videoID,
		Content: content,
	}
	commentInfo, err := rpc.CreateComment(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "CreateComment rpc failed, appUserID: %v, videoID: %v, content: %s", appUserID, videoID, content)
	}
	author, err := rpc.GetUser(ctx, &userproto.GetUserReq{
		AppUserId: appUserID,
		UserId:    commentInfo.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, commentInfo.UserId)
	}
	return toCommentDTO(commentInfo, toUserDTO(author)), nil
}

// DeleteComment
// delete a comment
func (c CommentAppService) DeleteComment(ctx context.Context, commentID int64, videoID int64) (err error) {
	//panic("implement me")
	err = rpc.DeleteComment(ctx, &commentproto.DeleteCommentReq{CommentId: commentID, VideoId: videoID})
	if err != nil {
		return errors.Wrapf(err, "DeleteComment rpc failed, commentID: %v", commentID)
	}
	return nil
}

// GetCommentList
// get comment list by video's id
func (c CommentAppService) GetCommentList(ctx context.Context, appUserID int64, videoID int64) (commentList []*pojo.Comment, err error) {
	//panic("implement me")
	commentInfos, err := rpc.GetCommentList(ctx, &commentproto.GetCommentListReq{VideoId: videoID})
	if err != nil {
		return nil, errors.Wrapf(err, "GetCommentList rpc failed, appUserID: %v, videoID: %v", appUserID, videoID)
	}

	n := len(commentInfos)
	authors := make([]*pojo.User, n)
	for i := 0; i < n; i++ {
		authorInfo, err := rpc.GetUser(ctx, &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    commentInfos[i].UserId, //获取评论的作者id
		})
		if err != nil {
			return nil, errors.Wrapf(err, "GetUser rpc failed, appUserID: %v, userId: %v", appUserID, commentInfos[i].UserId)
		}
		authors[i] = toUserDTO(authorInfo)
	}
	return toCommentDTOs(commentInfos, authors), nil
}

func toCommentDTO(commentInfo *commentproto.CommentInfo, user *pojo.User) *pojo.Comment {
	if commentInfo == nil {
		return nil
	}
	return &pojo.Comment{
		ID:         commentInfo.CommentId,
		User:       user,
		Content:    commentInfo.Content,
		CreateDate: commentInfo.CreateDate,
	}
}

func toCommentDTOs(commentInfos []*commentproto.CommentInfo, authors []*pojo.User) []*pojo.Comment {
	n := len(commentInfos)
	comments := make([]*pojo.Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = &pojo.Comment{
			ID:         commentInfos[i].CommentId,
			User:       authors[i],
			Content:    commentInfos[i].Content,
			CreateDate: commentInfos[i].CreateDate,
		}
	}
	return comments
}
