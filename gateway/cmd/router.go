package cmd

import (
	"github.com/gin-gonic/gin"
	"runedance/gateway/api"
	"runedance/gateway/api/authorization"
)

func register(r *gin.Engine) {
	unAuthGroup := r.Group("/douyin")
	{
		unAuthGroup.GET("/feed", api.Feed)
		unAuthGroup.POST("/user/register/", api.Create)
		unAuthGroup.POST("/user/login/", api.Check)
	}

	authGroup := r.Group("/douyin")
	authGroup.Use(authorization.JWT())
	{
		authGroup.POST("/publish/action/", api.Upload)
		authGroup.GET("/publish/list/", api.List)
		authGroup.POST("/favorite/action/", api.LikeAction)
		authGroup.GET("/favorite/list/", api.LikeList)
		authGroup.POST("/comment/action/", api.CommentAction)
		authGroup.GET("/comment/list/", api.CommentList)
		authGroup.GET("/user/", api.GetUserInfo)
		authGroup.POST("/relation/action/", api.Follow)
		authGroup.GET("/relation/follow/list/", api.FollowList)
		authGroup.GET("/relation/follower/list/", api.FanList)
		authGroup.GET("/relation/friend/list/", api.FriendList)
		authGroup.GET("/message/chat/", api.GetMessageList)
		authGroup.POST("/message/action/", api.HandleMessageAction)
	}
}
