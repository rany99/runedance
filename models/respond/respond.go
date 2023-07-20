package respond

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runedance/common/errorCode"
	"runedance/pkg/errStatu"
)

type BaseResp struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

var Success = BaseResp{
	Code: 0,
	Msg:  "success",
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, BaseResp{
		Code: int64(errorCode.HTTPCoder(errStatu.Code(err))),
		Msg:  err.Error(),
	})
}

func Send(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Success)
}
