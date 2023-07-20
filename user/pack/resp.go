package pack

import (
	"runedance/kitexGen/kitex_gen/userproto"
	"runedance/pkg/errorCode"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userproto.BaseResp {
	return baseResp(errorCode.ConvertErr(err))
}

func baseResp(err errorCode.ErrNo) *userproto.BaseResp {
	return &userproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
