// Code generated by hertz generator.

package douyin

import (
	"context"
	"runedance/common/constant"
	auth "runedance/gateway/api/authorization"
	"runedance/gateway/application"
	douyin "runedance/gatewayHertz/biz/model/douyin"
	"runedance/models/pojo"
	"runedance/models/respond"
	"runedance/pkg/errorCode"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	appUserID := c.GetInt64(constant.IdentityKey)
	switch req.ActionType {
	case constant.FollowUser:
		err := application.UserAppIns.FollowUser(ctx, appUserID, req.ToUserID)
		if err != nil {
			c.JSON(errorCode.ServiceErrCode, new(douyin.RelationActionResponse))
			return
		}
		c.JSON(consts.StatusOK, new(douyin.RelationActionResponse))
	case constant.UnFollowUser:
		err := application.UserAppIns.UnFollowUser(ctx, appUserID, req.ToUserID)
		if err != nil {
			c.JSON(errorCode.ServiceErrCode, new(douyin.RelationActionResponse))
			return
		}
		c.JSON(consts.StatusOK, new(douyin.RelationActionResponse))
	default:
		c.JSON(errorCode.ServiceErrCode, new(douyin.RelationActionResponse))
	}

}

// GetFollowList .
// @router /douyin/relation/follow/list/ [GET]
func GetFollowList(ctx context.Context, c *app.RequestContext) {

	var err error
	var req douyin.GetFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	appUserID := c.GetInt64(constant.IdentityKey)

	users, err := application.UserAppIns.GetFollowList(ctx, appUserID, req.GetUserID())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFollowListResponse))
		return
	}
	resp := &pojo.FollowListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	c.JSON(consts.StatusOK, resp)
}

// GetFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func GetFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	appUserID := c.GetInt64(constant.IdentityKey)
	users, err := application.UserAppIns.GetFanList(ctx, appUserID, req.GetUserID())

	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFollowerListResponse))
		return
	}
	resp := &pojo.FanListResp{
		BaseResp: respond.Success,
		UserList: users,
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFriendList .
// @router /douyin/relation/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//resp := new(douyin.GetFriendListResponse)
	appUserID := c.GetInt64(constant.IdentityKey)
	users, err := application.UserAppIns.GetFriendList(ctx, appUserID, req.GetUserID())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.FriendListResp{
		BaseResp: respond.Success,
		UserList: users,
	}
	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	appUserID := c.GetInt64(constant.IdentityKey)
	err = application.MessageAppIns.CreateMessage(ctx, appUserID, req.GetToUserID(), req.GetContent())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}

	c.JSON(consts.StatusOK, new(douyin.MessageActionResponse))
}

// GetMessageChat .
// @router /douyin/message/chat/ [POST]
func GetMessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	appUserID := c.GetInt64(constant.IdentityKey)
	messageList, err := application.MessageAppIns.GetMessageList(ctx, appUserID, req.GetToUserID())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := pojo.MessageListResp{
		BaseResp:    respond.Success,
		MessageList: messageList,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userId, err := application.UserAppIns.CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	token, err := auth.GenerateToken(userId)
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.UserRegisterResp{
		BaseResp: respond.Success,
		UserID:   userId,
		Token:    token,
	}
	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userId, err := application.UserAppIns.CheckUser(ctx, req.Username, req.Password)
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	token, err := auth.GenerateToken(userId)
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.UserLoginResp{
		BaseResp: respond.Success,
		UserID:   userId,
		Token:    token,
	}
	c.JSON(consts.StatusOK, resp)
}

// GetUser .
// @router /douyin/user/ [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	appUserID := c.GetInt64(constant.IdentityKey)
	user, err := application.UserAppIns.GetUser(ctx, appUserID, req.GetUserID())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.UserQueryResp{
		BaseResp: respond.Success,
		User:     user,
	}
	c.JSON(consts.StatusOK, resp)
}

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	appUserID := c.GetInt64(constant.IdentityKey)
	req := new(pojo.VideoFeedReq)
	err := c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	if req.LatestTime <= 0 {
		req.LatestTime = time.Now().Unix()
	}
	if req.LatestTime > time.Now().Unix() {
		req.LatestTime = time.Now().Unix()
	}
	videoList, nextTime, err := application.VideoAppIns.Feed(ctx, appUserID, req.LatestTime)
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.VideoFeedResp{
		BaseResp:  respond.Success,
		NextTime:  nextTime,
		VideoList: videoList,
	}
	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	appUserID := c.GetInt64(constant.IdentityKey)
	var req douyin.DouyinPublishActionRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	fileHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	if err := application.VideoAppIns.PublishVideo(ctx, appUserID, req.Title, fileHeader); err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	c.JSON(200, new(douyin.DouyinPublishActionResponse))
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	appUserID := c.GetInt64(constant.IdentityKey)
	var err error
	var req douyin.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	videoList, err := application.VideoAppIns.GetVideoList(ctx, appUserID, req.GetUserID())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.VideoListResp{
		BaseResp:  respond.Success,
		VideoList: videoList,
	}
	c.JSON(200, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	appUserID := c.GetInt64(constant.IdentityKey)
	var err error
	var req douyin.FavoriteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	switch req.ActionType {
	case constant.LikeVideo: // 点赞
		if err := application.VideoAppIns.LikeVideo(ctx, appUserID, req.GetVideoID()); err != nil {
			c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
			return
		} else {
			c.JSON(consts.StatusOK, new(douyin.FavoriteResponse))
			return
		}
	case constant.UnLikeVideo: // 取消点赞
		if err := application.VideoAppIns.UnLikeVideo(ctx, appUserID, req.GetVideoID()); err != nil {
			c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
			return
		} else {
			c.JSON(consts.StatusOK, new(douyin.FavoriteResponse))
			return
		}
	default:
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
	}
}

// GetFavoriteList .
// @router /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	appUserID := c.GetInt64(constant.IdentityKey)
	var err error
	var req douyin.GetFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	videos, err := application.VideoAppIns.GetLikeVideoList(ctx, appUserID, req.GetUserID())
	if err != nil {
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
		return
	}
	resp := &pojo.LikeListResp{
		BaseResp:  respond.Success,
		VideoList: videos,
	}
	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {

	appUserID := c.GetInt64(constant.IdentityKey)
	var err error
	var req douyin.CommentRequest
	//fmt.Println(c)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	switch req.ActionType {
	case constant.CreateComment: // create comment on a video
		comment, err := application.CommentAppIns.CreateComment(ctx, appUserID, req.GetVideoID(), req.GetCommentText())
		if err != nil {
			c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
			return
		}
		resp := &pojo.CreateCommentResp{
			BaseResp: respond.Success,
			Comment:  comment,
		}
		c.JSON(consts.StatusOK, resp)
		return
	case constant.DeleteComment: // delete one comment
		if err := application.CommentAppIns.DeleteComment(ctx, req.GetVideoID(), req.GetVideoID()); err != nil {
			c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
			return
		}
		c.JSON(consts.StatusOK, new(douyin.CommentResponse))
		return
	default:
		c.JSON(errorCode.ServiceErrCode, new(douyin.GetFriendListResponse))
	}
}

// GetCommentList .
// @router /douyin/comment/list/ [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	appUserID := c.GetInt64(constant.IdentityKey)
	var err error
	var req douyin.GetCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	comments, err := application.CommentAppIns.GetCommentList(ctx, appUserID, req.GetVideoID())
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := &pojo.CommentListResp{
		BaseResp:    respond.Success,
		CommentList: comments,
	}
	c.JSON(consts.StatusOK, resp)
}
