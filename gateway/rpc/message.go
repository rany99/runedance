package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"runedance/common/config"
	"runedance/common/constant"
	"runedance/kitexGen/kitex_gen/messageproto"
	"runedance/kitexGen/kitex_gen/messageproto/messageservice"
	errno "runedance/pkg/errStatu"
	"runedance/pkg/middleware"
	"time"
)

var messageClient messageservice.Client

func initMessageRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Server.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		constant.MessageDomainServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(time.Minute),                // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

func CreateMessage(ctx context.Context, req *messageproto.CreateMessageReq) error {
	resp, err := messageClient.CreateMessage(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetMessageList(ctx context.Context, req *messageproto.GetMessageListReq) ([]*messageproto.MessageInfo, error) {
	resp, err := messageClient.GetMessageList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.New(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.MessageInfos, nil
}
