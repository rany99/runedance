package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/comment/dao/dal"
	"runedance/comment/dao/redis"
	"runedance/comment/pulsar"
	"runedance/common/config"
	"runedance/kitexGen/kitex_gen/commentproto"
	"testing"
)

func initCreateCommentTest() {
	config.InitConfig()
	dal.Init()
	redis.Init()
	pulsar.Init()
}

func TestCreateCommentService(t *testing.T) {
	initCreateCommentTest()
	ctx := context.Background()
	req := &commentproto.CreateCommentReq{
		UserId:  15,
		VideoId: 14,
		Content: "test comment",
	}
	comments, err := NewCreateCommentService(ctx).CreateComment(req.UserId, req.VideoId, req.Content)
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println(comments)
}
