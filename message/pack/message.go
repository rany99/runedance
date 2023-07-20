package pack

import (
	"runedance/kitexGen/kitex_gen/messageproto"
	"runedance/message/dao/dal/entity"
	redisModel "runedance/message/dao/redis/model"
)

func Message(message *entity.Message) *messageproto.MessageInfo {
	return &messageproto.MessageInfo{
		MessageId:  message.MessageUUID,
		FromUserId: message.FromUserId,
		ToUserId:   message.ToUserId,
		Content:    message.Contents,
		CreateTime: message.CreateTime,
	}
}

func Messages(messages []*entity.Message) []*messageproto.MessageInfo {
	messageInfos := make([]*messageproto.MessageInfo, len(messages))
	for i, message := range messages {
		messageInfos[i] = Message(message)
	}
	return messageInfos
}

func MessageFromRedisModel(message *redisModel.MessageRedis) *entity.Message {
	return &entity.Message{
		FromUserId:  message.FromUserId,
		ToUserId:    message.ToUserId,
		Contents:    message.Content,
		MessageUUID: message.MessageId,
		CreateTime:  message.CreateTime,
	}
}
