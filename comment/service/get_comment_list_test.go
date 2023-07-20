package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/comment/dao/dal"
	"runedance/comment/dao/redis"
	"runedance/common/config"
	"runedance/kitexGen/kitex_gen/commentproto"
	"testing"
)

func initGetCommentListTest() {
	config.InitConfig()
	dal.Init()
	redis.Init()
}

func TestGetCommentListService(t *testing.T) {
	initGetCommentListTest()
	ctx := context.Background()
	req := &commentproto.GetCommentListReq{VideoId: 3}
	comments, err := NewGetCommentListService(ctx).GetCommentList(req.VideoId)
	if err != nil {
		klog.Error(err.Error())
	}
	fmt.Println(comments)
}
