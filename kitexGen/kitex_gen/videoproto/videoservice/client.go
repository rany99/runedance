// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	videoproto "runedance/kitexGen/kitex_gen/videoproto"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateVideo(ctx context.Context, Req *videoproto.CreateVideoReq, callOptions ...callopt.Option) (r *videoproto.CreateVideoResp, err error)
	GetVideoListByUserId(ctx context.Context, Req *videoproto.GetVideoListByUserIdReq, callOptions ...callopt.Option) (r *videoproto.GetVideoListByUserIdResp, err error)
	GetVideoListByTime(ctx context.Context, Req *videoproto.GetVideoListByTimeReq, callOptions ...callopt.Option) (r *videoproto.GetVideoListByTimeResp, err error)
	LikeVideo(ctx context.Context, Req *videoproto.LikeVideoReq, callOptions ...callopt.Option) (r *videoproto.LikeVideoResp, err error)
	UnLikeVideo(ctx context.Context, Req *videoproto.UnLikeVideoReq, callOptions ...callopt.Option) (r *videoproto.UnLikeVideoResp, err error)
	GetLikeVideoList(ctx context.Context, Req *videoproto.GetLikeVideoListReq, callOptions ...callopt.Option) (r *videoproto.GetLikeVideoListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) CreateVideo(ctx context.Context, Req *videoproto.CreateVideoReq, callOptions ...callopt.Option) (r *videoproto.CreateVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVideo(ctx, Req)
}

func (p *kVideoServiceClient) GetVideoListByUserId(ctx context.Context, Req *videoproto.GetVideoListByUserIdReq, callOptions ...callopt.Option) (r *videoproto.GetVideoListByUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoListByUserId(ctx, Req)
}

func (p *kVideoServiceClient) GetVideoListByTime(ctx context.Context, Req *videoproto.GetVideoListByTimeReq, callOptions ...callopt.Option) (r *videoproto.GetVideoListByTimeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoListByTime(ctx, Req)
}

func (p *kVideoServiceClient) LikeVideo(ctx context.Context, Req *videoproto.LikeVideoReq, callOptions ...callopt.Option) (r *videoproto.LikeVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeVideo(ctx, Req)
}

func (p *kVideoServiceClient) UnLikeVideo(ctx context.Context, Req *videoproto.UnLikeVideoReq, callOptions ...callopt.Option) (r *videoproto.UnLikeVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnLikeVideo(ctx, Req)
}

func (p *kVideoServiceClient) GetLikeVideoList(ctx context.Context, Req *videoproto.GetLikeVideoListReq, callOptions ...callopt.Option) (r *videoproto.GetLikeVideoListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikeVideoList(ctx, Req)
}