package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/common/constant"
	"runedance/video/dao/dal"
	"runedance/video/dao/redis"
)

func LikeVideoConsume(ctx context.Context, client pulsar.Client) error {
	//listen the channel
	channel := make(chan pulsar.ConsumerMessage, 100)
	var likeVideoJS LikeVideoJSON
	consumerJS := pulsar.NewJSONSchema(LikeVideoSchemaDef, nil)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            constant.LikeVideoTopic,
		SubscriptionName: "sub-1",
		Schema:           consumerJS,
		Type:             pulsar.Shared,
		MessageChannel:   channel,
	})
	if err != nil {
		return err
	}
	defer consumer.Close()

	for cm := range channel {
		consumer := cm.Consumer
		msg := cm.Message
		err = msg.GetSchemaValue(&likeVideoJS)
		if err != nil {
			klog.Error(err)
		}
		err = consumer.Ack(msg)
		if err != nil {
			klog.Error(err)
		}
		switch likeVideoJS.ActionType {
		case constant.LikeVideo:
			if err := dal.LikeVideo(ctx, likeVideoJS.UserID, likeVideoJS.VideoID); err != nil {
				klog.Error("mysql error:", err)
				redis.DeleteLikeKey(likeVideoJS.UserID, likeVideoJS.VideoID)
				if err != nil {
					klog.Error("del redis key err", err)
				}
			}
			break
		case constant.UnLikeVideo:
			if err := dal.UnLikeVideo(ctx, likeVideoJS.UserID, likeVideoJS.VideoID); err != nil {
				klog.Error("mysql error:", err)
				redis.DeleteLikeKey(likeVideoJS.UserID, likeVideoJS.VideoID)
				if err != nil {
					klog.Error("del redis key err", err)
				}
			}
			break
		}

	}

	return nil
}
