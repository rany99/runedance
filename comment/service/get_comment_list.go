package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/go-multierror"
	"runedance/comment/dao/dal"
	"runedance/comment/dao/redis"
	"runedance/comment/pack"
	"runedance/kitexGen/kitex_gen/commentproto"
)

type GetCommentListService struct {
	ctx context.Context
}

func NewGetCommentListService(ctx context.Context) *GetCommentListService {
	return &GetCommentListService{
		ctx: ctx,
	}
}

func (s *GetCommentListService) GetCommentList(videoId int64) ([]*commentproto.CommentInfo, error) {
	comments, redisErr := redis.GetCommentList(videoId)
	//读取Redis和DB有出错的时候的一致性控制
	if redisErr != nil {
		// 从redis中读取评论列表失败，转而从DB中读
		klog.Error("GetCommentList from redis failed, " + redisErr.Error() + ", getting from DB..")
		commentsDB, dbErr := dal.GetCommentList(s.ctx, videoId)
		if dbErr != nil {
			// 完蛋，数据库和缓存全都读失败了，抛出合并的error
			klog.Error("DB and Redis GetCommentList both failed, " + dbErr.Error())
			return nil, multierror.Append(redisErr, dbErr)
		}
		// redis失败，db成功
		// 需要刷新redis缓存，将db中读取的写入redis
		// 刷新缓存不用担心并发控制
		err := redis.AddCommentList(commentsDB)
		if err != nil {
			klog.Error("redis refresh failed, " + err.Error())
		}
		klog.Info("DB GetCommentList succeed! ")
		return pack.Comments(commentsDB), nil
	}
	return pack.RedisComments(comments), nil
}
