package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/common/util"
	"runedance/message/dao/dal/entity"
)

func CreateMessage(ctx context.Context, userId int64, toUserID int64, content string, createTime int64) error {
	uuid, err := util.GenSnowFlake(0)
	if err != nil {
		klog.Error("Failed to generate UUID" + err.Error())
		return err
	}

	message := entity.Message{
		FromUserId:  userId,
		ToUserId:    toUserID,
		Contents:    content,
		MessageUUID: int64(uuid),
		CreateTime:  createTime,
	}
	err = DB.WithContext(ctx).Create(&message).Error
	if err != nil {
		klog.Error("create message fail " + err.Error())
		return err
	}
	return nil
}

func GetMessageList(ctx context.Context, userId int64, toUserID int64, latestTime int64) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ? AND create_time >= ?",
		userId, toUserID, latestTime).Or("from_user_id = ? AND to_user_id = ? AND create_time >= ?",
		toUserID, userId, latestTime).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
