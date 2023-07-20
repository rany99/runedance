package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"runedance/common/constant"
	"runedance/kitexGen/kitex_gen/userproto"
	"runedance/user/dao/dal"
	"runedance/user/dao/redis"
	"runedance/user/pulsar"
)

type FollowUserService struct {
	ctx context.Context
}

// NewFollowUserService new FollowUserService
func NewFollowUserService(ctx context.Context) *FollowUserService {
	return &FollowUserService{
		ctx: ctx,
	}
}

// FollowUser Follow user by id
func (s *FollowUserService) FollowUser(req *userproto.FollowUserReq) error {
	if req.FanUserId == req.FollowedUserId {
		return errors.New("can't follow yourself")
	}
	userId := req.FanUserId
	followId := req.FollowedUserId
	if exist, _ := redis.IsFollowKeyExist(userId); exist == false {
		followIdDalList, err := dal.GetFollowList(s.ctx, userId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		redis.AddFollowList(userId, followIdDalList)
	}
	if exist, _ := redis.IsFanKeyExist(followId); exist == false {
		lock := redis.NewUserKeyLock(userId, constant.FanRedisPrefix)
		err := lock.Lock(s.ctx)
		if err != nil {
			klog.Error(err)
		}
		// DCL
		if existDC, _ := redis.IsFollowKeyExist(userId); existDC == false {
			fanIdDalList, err := dal.GetFollowList(s.ctx, followId)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			redis.AddFollowList(followId, fanIdDalList)
		}
		lock.Unlock()
	}
	err := redis.AddRelation(userId, followId)
	if err != nil {
		return err
	}

	if err := pulsar.FollowUserProduce(s.ctx, userId, followId); err != nil {
		return err
	}
	redis.AddBloomKey(constant.FollowRedisPrefix, userId)
	redis.AddBloomKey(constant.FanRedisPrefix, followId)
	return nil
}
