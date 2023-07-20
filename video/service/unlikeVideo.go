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

type UnLikeVideoService struct {
	ctx context.Context
}

func NewUnLikeVideoService(ctx context.Context) *UnLikeVideoService {
	return &UnLikeVideoService{ctx: ctx}
}

func (s *UnLikeVideoService) UnLikeVideo(req *videoproto.UnLikeVideoReq) error {
	span := Tracer.StartSpan("unlike_video")
	defer span.Finish()
	s.ctx = opentracing.ContextWithSpan(s.ctx, span)
	userId := req.UserId
	videoID := req.VideoId
	isLikeKeyExist, err := redis.IsLikeKeyExist(userId)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == true {
		// 如果redis有这个userId的记录，则需要在redis中删去这条like记录，确保和mysql一致
		isLikeById, err := redis.GetIsLikeById(userId, videoID)
		if err != nil {
			klog.Error(err)
		}
		if isLikeById == false {
			return nil
		}
		if err := redis.DeleteLike(userId, videoID); err != nil {
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
		if isLikeById == false {
			return nil
		}
		if err := redis.DeleteLike(userId, videoID); err != nil {
			klog.Error(err)
		}
	}
	if err := pulsar.UnLikeVideoProduce(s.ctx, userId, videoID); err != nil {
		return err
	}
	return nil
}
