package api

import (
	"github.com/gin-gonic/gin"
	"runedance/common/constant"
	auth "runedance/gateway/api/authorization"
	"runedance/gateway/application"
	"runedance/models/pojo"
	"runedance/models/respond"
	"time"
)

func Feed(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil { // Cases in which the user is not logged in
		appUserID = -1
	}

	param := new(pojo.VideoFeedReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	if param.LatestTime <= 0 {
		param.LatestTime = time.Now().Unix()
	}
	if param.LatestTime > time.Now().Unix() {
		param.LatestTime = time.Now().Unix()
	}

	videoList, nextTime, err := application.VideoAppIns.Feed(c, appUserID, param.LatestTime)
	if err != nil {
		respond.Error(c, err)
		return
	}

	resp := &pojo.VideoFeedResp{
		BaseResp:  respond.Success,
		NextTime:  nextTime,
		VideoList: videoList,
	}
	respond.Send(c, resp)
}

// Upload upload video (POST)
// Login user to select video upload
func Upload(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.VideoUploadReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	fileHeader, err := c.FormFile("data")
	if err != nil {
		respond.Error(c, err)
		return
	}
	if err := application.VideoAppIns.PublishVideo(c, appUserID, param.Title, fileHeader); err != nil {
		respond.Error(c, err)
		return
	}
	respond.OK(c)
}

// List upload list (GET)
// Log in to the user's video posting list and directly list all the videos that the user has contributed to
func List(c *gin.Context) {
	appUserID := c.GetInt64(constant.IdentityKey)
	param := new(pojo.VideoListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}

	videoList, err := application.VideoAppIns.GetVideoList(c, appUserID, param.UserId)
	if err != nil {
		respond.Error(c, err)
		return
	}

	resp := &pojo.VideoListResp{
		BaseResp:  respond.Success,
		VideoList: videoList,
	}
	respond.Send(c, resp)
}
