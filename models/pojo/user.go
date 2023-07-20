package pojo

import (
	"runedance/models/respond"
)

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow"` // true-已关注，false-未关注
}

// 获取用户信息
type UserQueryReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}

type UserQueryResp struct {
	respond.BaseResp
	User *User `json:"user,omitempty"`
}

// 用户注册
type UserRegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32" msg:"最长32个字符串"`
	Password string `form:"password" json:"password" binding:"required,max=32" msg:"最长32个字符串"`
}

type UserRegisterResp struct {
	respond.BaseResp
	UserID int64  `json:"user_id"`
	Token  string `json:"token"` // 用户鉴权token
}

// 用户登录
type UserLoginReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

type UserLoginResp struct {
	respond.BaseResp
	UserID int64  `json:"user_id"`
	Token  string `json:"token"` // 用户鉴权token
}

// 关注操作
type FollowOperationReq struct {
	ToUserId   int64 `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int   `form:"action_type" json:"action_type" binding:"required" msg:"1-关注，2-取消关注"`
}

// 关注列表
type FollowListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}

type FollowListResp struct {
	respond.BaseResp
	UserList []*User `json:"user_list,omitempty"`
}

// 粉丝列表
type FanListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}

type FanListResp struct {
	respond.BaseResp
	UserList []*User `json:"user_list,omitempty"`
}

// 好友列表
type FriendListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}

type FriendListResp struct {
	respond.BaseResp
	UserList []*User `json:"user_list,omitempty"`
}
