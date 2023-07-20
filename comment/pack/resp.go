package pack

import (
	"runedance/kitexGen/kitex_gen/commentproto"
	"runedance/pkg/errorCode"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *commentproto.BaseResp {
	return baseResp(errorCode.ConvertErr(err))
}

func baseResp(err errorCode.ErrNo) *commentproto.BaseResp {
	return &commentproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
