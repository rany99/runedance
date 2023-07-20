package api

import (
	"github.com/gin-gonic/gin"
	"runedance/common/constant"
	"runedance/gateway/application"
	"runedance/models/pojo"
	"runedance/models/respond"
	"runedance/pkg/errorCode"
)

func Follow(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.FollowOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case constant.FollowUser:
		err := application.UserAppIns.FollowUser(c, appUserID, param.ToUserId)
		if err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	case constant.UnFollowUser:
		err := application.UserAppIns.UnFollowUser(c, appUserID, param.ToUserId)
		if err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, errorCode.ParamErr)
	}
}

func FollowList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.FollowListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	users, err := application.UserAppIns.GetFollowList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.FollowListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	respond.Send(c, resp)
}

func FanList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.FanListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	users, err := application.UserAppIns.GetFanList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.FanListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	respond.Send(c, resp)
}

func FriendList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.FriendListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	users, err := application.UserAppIns.GetFriendList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.FriendListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	respond.Send(c, resp)
}
