package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/opentracing/opentracing-go"
	"runedance/kitexGen/kitex_gen/videoproto"
	"runedance/video/dao/dal"
	"runedance/video/dao/redis"
	"runedance/video/pulsar"
)

type LikeVideoService struct {
	ctx context.Context
}

func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

func (s *LikeVideoService) LikeVideo(req *videoproto.LikeVideoReq) error {
	span := Tracer.StartSpan("like_video")
	defer span.Finish()
	s.ctx = opentracing.ContextWithSpan(s.ctx, span)
	userId := req.UserId
	videoID := req.VideoId
	isLikeKeyExist, err := redis.IsLikeKeyExist(userId)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == true {
		// 如果redis有这个userId的记录，则需要在redis中再加入这条新的点赞的操作，确保和mysql一致
		isLikeById, err := redis.GetIsLikeById(userId, videoID)
		if err != nil {
			klog.Error(err)
		}
		if isLikeById == true {
			return nil
		}
		if err := redis.AddLike(userId, videoID); err != nil {
			klog.Error(err)
		}
	} else {
		// 如果redis没有这个userId的记录，则去mysql查询一次点赞列表进行缓存
		likeList, err := dal.MGetLikeList(s.ctx, userId)
		if err != nil {
			klog.Error(err)
		}
		if err := redis.AddLikeList(userId, likeList); err != nil {
			klog.Error(err)
			return err
		}
		isLikeById, err := redis.GetIsLikeById(userId, videoID)
		if err != nil {
			klog.Error(err)
		}
		if isLikeById == true {
			return nil
		}
		if err := redis.AddLike(userId, videoID); err != nil {
			klog.Error(err)
		}
	}
	if err = pulsar.LikeVideoProduce(s.ctx, userId, videoID); err != nil {
		return err
	}
	return nil
}
