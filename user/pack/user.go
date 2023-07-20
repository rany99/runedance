package pack

import (
	"runedance/kitexGen/kitex_gen/userproto"
	"runedance/user/dao/dal/entity"
	redisModel "runedance/user/dao/redis/model"
)

func PackUserRedis(userRedis *redisModel.UserRedis) *userproto.UserInfo {
	if userRedis == nil {
		return nil
	}
	return &userproto.UserInfo{
		UserId:        userRedis.UserId,
		Username:      userRedis.UserName,
		FollowCount:   userRedis.FollowCnt,
		FollowerCount: userRedis.FanCnt,
		WorkCount:     userRedis.WorkCnt,
		FavoriteCount: userRedis.FavoriteCnt,
	}
}

func PackUserDal(user *entity.User) *userproto.UserInfo {
	if user == nil {
		return nil
	}
	return &userproto.UserInfo{
		UserId:        int64(user.ID),
		Username:      user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		WorkCount:     user.WorkCount,
		FavoriteCount: user.FavoriteCount,
	}
}
