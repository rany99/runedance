package errorCode

import (
	"net/http"
	"runedance/pkg/errStatu"
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
	UnauthorizedErrCode     = 10006
)

var (
	Success             = errStatu.New(SuccessCode, "Success")
	ServiceErr          = errStatu.New(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = errStatu.New(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = errStatu.New(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = errStatu.New(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = errStatu.New(UserAlreadyExistErrCode, "User already exists")
	UnauthorizedErr     = errStatu.New(UnauthorizedErrCode, "JWT Token Unauthorized")
)

var mapper = map[int64]int{
	SuccessCode:             http.StatusOK,
	ServiceErrCode:          http.StatusInternalServerError,
	ParamErrCode:            http.StatusBadRequest,
	LoginErrCode:            http.StatusBadRequest,
	UserNotExistErrCode:     http.StatusBadRequest,
	UserAlreadyExistErrCode: http.StatusBadRequest,
	UnauthorizedErrCode:     http.StatusUnauthorized,
}

func HTTPCoder(code int64) int {
	if http, ok := mapper[code]; ok {
		return http
	}
	return http.StatusInternalServerError
}
