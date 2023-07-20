// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	messageproto "runedance/kitexGen/kitex_gen/messageproto"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

var messageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageService"
	handlerType := (*messageproto.MessageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateMessage":  kitex.NewMethodInfo(createMessageHandler, newCreateMessageArgs, newCreateMessageResult, false),
		"GetMessageList": kitex.NewMethodInfo(getMessageListHandler, newGetMessageListArgs, newGetMessageListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "message",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(messageproto.CreateMessageReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(messageproto.MessageService).CreateMessage(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateMessageArgs:
		success, err := handler.(messageproto.MessageService).CreateMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateMessageResult)
		realResult.Success = success
	}
	return nil
}
func newCreateMessageArgs() interface{} {
	return &CreateMessageArgs{}
}

func newCreateMessageResult() interface{} {
	return &CreateMessageResult{}
}

type CreateMessageArgs struct {
	Req *messageproto.CreateMessageReq
}

func (p *CreateMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(messageproto.CreateMessageReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateMessageArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateMessageArgs) Unmarshal(in []byte) error {
	msg := new(messageproto.CreateMessageReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateMessageArgs_Req_DEFAULT *messageproto.CreateMessageReq

func (p *CreateMessageArgs) GetReq() *messageproto.CreateMessageReq {
	if !p.IsSetReq() {
		return CreateMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateMessageResult struct {
	Success *messageproto.CreateMessageResp
}

var CreateMessageResult_Success_DEFAULT *messageproto.CreateMessageResp

func (p *CreateMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(messageproto.CreateMessageResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateMessageResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateMessageResult) Unmarshal(in []byte) error {
	msg := new(messageproto.CreateMessageResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateMessageResult) GetSuccess() *messageproto.CreateMessageResp {
	if !p.IsSetSuccess() {
		return CreateMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*messageproto.CreateMessageResp)
}

func (p *CreateMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getMessageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(messageproto.GetMessageListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(messageproto.MessageService).GetMessageList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetMessageListArgs:
		success, err := handler.(messageproto.MessageService).GetMessageList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetMessageListResult)
		realResult.Success = success
	}
	return nil
}
func newGetMessageListArgs() interface{} {
	return &GetMessageListArgs{}
}

func newGetMessageListResult() interface{} {
	return &GetMessageListResult{}
}

type GetMessageListArgs struct {
	Req *messageproto.GetMessageListReq
}

func (p *GetMessageListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(messageproto.GetMessageListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetMessageListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetMessageListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetMessageListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMessageListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetMessageListArgs) Unmarshal(in []byte) error {
	msg := new(messageproto.GetMessageListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetMessageListArgs_Req_DEFAULT *messageproto.GetMessageListReq

func (p *GetMessageListArgs) GetReq() *messageproto.GetMessageListReq {
	if !p.IsSetReq() {
		return GetMessageListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMessageListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMessageListResult struct {
	Success *messageproto.GetMessageListResp
}

var GetMessageListResult_Success_DEFAULT *messageproto.GetMessageListResp

func (p *GetMessageListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(messageproto.GetMessageListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetMessageListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetMessageListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetMessageListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMessageListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetMessageListResult) Unmarshal(in []byte) error {
	msg := new(messageproto.GetMessageListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMessageListResult) GetSuccess() *messageproto.GetMessageListResp {
	if !p.IsSetSuccess() {
		return GetMessageListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMessageListResult) SetSuccess(x interface{}) {
	p.Success = x.(*messageproto.GetMessageListResp)
}

func (p *GetMessageListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateMessage(ctx context.Context, Req *messageproto.CreateMessageReq) (r *messageproto.CreateMessageResp, err error) {
	var _args CreateMessageArgs
	_args.Req = Req
	var _result CreateMessageResult
	if err = p.c.Call(ctx, "CreateMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMessageList(ctx context.Context, Req *messageproto.GetMessageListReq) (r *messageproto.GetMessageListResp, err error) {
	var _args GetMessageListArgs
	_args.Req = Req
	var _result GetMessageListResult
	if err = p.c.Call(ctx, "GetMessageList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
