package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/common/util"
	"runedance/kitexGen/kitex_gen/messageproto"
	"runedance/message/dao/dal"
	"runedance/message/dao/redis"
	"runedance/message/dao/redis/model"
	"runedance/message/pulsar"
	"time"
)

type CreateMessageService struct {
	ctx context.Context
}

func NewCreateMessageService(ctx context.Context) *CreateMessageService {
	return &CreateMessageService{
		ctx: ctx,
	}
}

func (s *CreateMessageService) CreateMessage(req *messageproto.CreateMessageReq) error {
	exists, err := redis.IsMessageKeyExist(req.UserId, req.ToUserId)
	if err != nil {
		klog.Error(err)
	}

	if !exists {
		// fetch messagelist into cache if not exists
		messages, err := dal.GetMessageList(s.ctx, req.UserId, req.ToUserId, 0)
		if err != nil {
			return err
		}
		err = redis.AddMessageList(req.UserId, req.ToUserId, messages)
		if err != nil {
			return err
		}
	}

	uuid, err := util.GenSnowFlake(0)
	if err != nil {
		klog.Error("Failed to generate UUID" + err.Error())
		return err
	}

	createTime := time.Now().Unix()

	message := model.MessageRedis{
		FromUserId: req.UserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
		MessageId:  int64(uuid),
		CreateTime: createTime,
	}
	err = redis.AddMessage(req.UserId, req.ToUserId, message)
	err = redis.AddMessageLatestTime(req.UserId, req.ToUserId, createTime+1)
	if err := pulsar.CreateMessageProduce(s.ctx, req.UserId, req.ToUserId, req.Content, createTime); err != nil {
		return err
	}
	return err
}
