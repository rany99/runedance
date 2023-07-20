package pack

import (
	"runedance/kitexGen/kitex_gen/messageproto"
	"runedance/pkg/errorCode"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *messageproto.BaseResp {
	return baseResp(errorCode.ConvertErr(err))
}

func baseResp(err errorCode.ErrNo) *messageproto.BaseResp {
	return &messageproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
