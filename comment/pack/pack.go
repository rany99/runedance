package pack

import (
	"runedance/comment/dao/dal/entity"
	"runedance/kitexGen/kitex_gen/commentproto"
	"time"
)

func Comment(comment *entity.Comment) *commentproto.CommentInfo {
	return &commentproto.CommentInfo{
		CommentId:  comment.CommentUUID,
		UserId:     comment.UserId,
		Content:    comment.Contents,
		CreateDate: time.Unix(comment.CreateTime, 0).Format("01-02"),
	}
}

func Comments(comments []*entity.Comment) []*commentproto.CommentInfo {
	commentInfos := make([]*commentproto.CommentInfo, len(comments))
	for i, comment := range comments {
		commentInfos[i] = Comment(comment)
	}
	return commentInfos
}

func redisComment(comment redisModel.CommentRedis) *commentproto.CommentInfo {
	return &commentproto.CommentInfo{
		CommentId:  comment.CommentId,
		UserId:     comment.UserId,
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CreateTime, 0).Format("01-02"),
	}
}
func RedisComments(comments []redisModel.CommentRedis) []*commentproto.CommentInfo {
	commentInfos := make([]*commentproto.CommentInfo, len(comments))
	for i, comment := range comments {
		commentInfos[i] = redisComment(comment)
	}
	return commentInfos
}
