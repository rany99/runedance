package api

import (
	"github.com/gin-gonic/gin"
	"runedance/common/constant"
	"runedance/gateway/api/authorization"
	"runedance/gateway/application"
	"runedance/models/pojo"
	"runedance/models/respond"
)

func GetUserInfo(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.UserQueryReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	//调用app层接口
	user, err := application.UserAppIns.GetUser(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.UserQueryResp{
		BaseResp: respond.Success,
		User:     user,
	}
	respond.Send(c, resp)
}

// Create User registration interface
func Create(c *gin.Context) {
	param := new(pojo.UserRegisterReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	userId, err := application.UserAppIns.CreateUser(c, param.Username, param.Password)
	if err != nil {
		respond.Error(c, err)
		return
	}

	token, err := authorization.GenerateToken(userId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.UserRegisterResp{
		BaseResp: respond.Success,
		UserID:   userId,
		Token:    token,
	}
	respond.Send(c, resp)
}

func Check(c *gin.Context) {
	param := new(pojo.UserLoginReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	userId, err := application.UserAppIns.CheckUser(c, param.Username, param.Password)
	if err != nil {
		respond.Error(c, err)
		return
	}
	token, err := authorization.GenerateToken(userId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.UserLoginResp{
		BaseResp: respond.Success,
		UserID:   userId,
		Token:    token,
	}
	respond.Send(c, resp)
}
