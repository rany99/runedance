package api

import (
	"github.com/gin-gonic/gin"
	"runedance/common/constant"
	"runedance/gateway/application"
	"runedance/models/pojo"
	"runedance/models/respond"
	"runedance/pkg/errorCode"
)

// LikeAction (POST)
// Like and unlike operations for videos by logged-in users
func LikeAction(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.LikeOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case constant.LikeVideo: // 点赞
		if err := application.VideoAppIns.LikeVideo(c, appUserID, param.VideoId); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	case constant.UnLikeVideo: // 取消点赞
		if err := application.VideoAppIns.UnLikeVideo(c, appUserID, param.VideoId); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, errorCode.ParamErr)
	}
}

// LikeList (GET)
// All liked videos by logged-in users
func LikeList(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.LikeListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	videos, err := application.VideoAppIns.GetLikeVideoList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &pojo.LikeListResp{
		BaseResp:  respond.Success,
		VideoList: videos,
	}
	respond.Send(c, resp)
}
