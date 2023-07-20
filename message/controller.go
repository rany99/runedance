package message

import (
	"context"
	"runedance/kitexGen/kitex_gen/messageproto"
	"runedance/message/pack"
	"runedance/message/service"
	"runedance/pkg/errorCode"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// CreateMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) CreateMessage(ctx context.Context, req *messageproto.CreateMessageReq) (resp *messageproto.CreateMessageResp, err error) {
	resp = new(messageproto.CreateMessageResp)

	if req.UserId < 0 || req.ToUserId < 0 || len(req.Content) == 0 { // Empty messages are not allowed
		resp.BaseResp = pack.BuildBaseResp(errorCode.ParamErr)
		return resp, nil
	}

	err = service.NewCreateMessageService(ctx).CreateMessage(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errorCode.Success)
	return resp, nil
}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *messageproto.GetMessageListReq) (resp *messageproto.GetMessageListResp, err error) {
	resp = new(messageproto.GetMessageListResp)

	if req.UserId < 0 || req.ToUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errorCode.ParamErr)
		return resp, nil
	}
	messages, err := service.NewGetMessageListService(ctx).GetMessageList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errorCode.Success)
	resp.MessageInfos = messages
	return resp, nil
}
