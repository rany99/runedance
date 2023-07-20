package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/common/constant"
	"runedance/message/dao/dal"
	"runedance/message/dao/redis"
)

func CreateMessageConsume(ctx context.Context, client pulsar.Client) error {
	//listen the channel
	channel := make(chan pulsar.ConsumerMessage, 100)
	var createMessageJS CreateMessageJSON
	consumerJS := pulsar.NewJSONSchema(CreateMessageSchemaDef, nil)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            constant.CreateMessageTopic,
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
		err = msg.GetSchemaValue(&createMessageJS)
		if err != nil {
			klog.Error(err)
		}
		err = consumer.Ack(msg)
		if err != nil {
			klog.Error(err)
		}

		if err := dal.CreateMessage(ctx, createMessageJS.UserId, createMessageJS.ToUserId, createMessageJS.Content, createMessageJS.CreateTime); err != nil {
			klog.Error("mysql error:", err)
			err = redis.DeleteMessageKey(createMessageJS.UserId, createMessageJS.ToUserId)
			if err != nil {
				klog.Error("del redis key err", err)
			}
		}
	}

	return nil
}
