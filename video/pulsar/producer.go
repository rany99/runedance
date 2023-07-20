package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"runedance/common/constant"
)

func LikeVideoProduce(ctx context.Context, userId int64, videoId int64) error {
	_, err := p_like_video.Send(ctx, &pulsar.ProducerMessage{
		Value: &LikeVideoJSON{
			UserID:     userId,
			VideoID:    videoId,
			ActionType: constant.LikeVideo,
		},
	})
	return err
}

func UnLikeVideoProduce(ctx context.Context, userId int64, videoId int64) error {
	_, err := p_like_video.Send(ctx, &pulsar.ProducerMessage{
		Value: &LikeVideoJSON{
			UserID:     userId,
			VideoID:    videoId,
			ActionType: constant.UnLikeVideo,
		},
	})
	return err
}
