package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"runedance/comment/dao/dal"
	"runedance/comment/dao/redis"
	"runedance/common/config"
	"runedance/kitexGen/kitex_gen/commentproto"
	"testing"
)

func initDeleteCommentTest() {
	config.InitConfig()
	dal.Init()
	redis.Init()
}

func TestDeleteCommentService(t *testing.T) {
	initDeleteCommentTest()
	ctx := context.Background()
	req := &commentproto.DeleteCommentReq{
		CommentId: 447637918983389184,
		VideoId:   7,
	}
	err := NewDeleteCommentService(ctx).DeleteComment(req.CommentId, req.VideoId)
	if err != nil {
		klog.Error(err.Error())
	}
	//fmt.Println(comments)
}
