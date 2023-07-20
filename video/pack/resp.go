package pack

import (
	"runedance/kitexGen/kitex_gen/videoproto"
	"runedance/pkg/errorCode"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *videoproto.BaseResp {
	return baseResp(errorCode.ConvertErr(err))
}

func baseResp(err errorCode.ErrNo) *videoproto.BaseResp {
	return &videoproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
