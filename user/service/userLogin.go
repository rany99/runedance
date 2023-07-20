package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"runedance/kitexGen/kitex_gen/userproto"
	"runedance/pkg/errorCode"
	"runedance/user/dao/dal"
	"runedance/user/dao/redis"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *userproto.CheckUserReq) (userId int64, err error) {
	username := req.UserAccount.Username
	//Verify that the user is locked out
	lock, _ := redis.IsLock(username)

	if lock {
		//Get Expiration Time
		min, _ := redis.GetUnlockTime(username)
		return 0, errorCode.NewLoginFailedTooManyErr(min)
	}

	//Check if the user entered the password correctly
	user, err := dal.GetUserByName(s.ctx, req.UserAccount.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return 0, errorCode.LoginErr
		}
		return 0, err
	}
	//param1 hashedPassword(stored) param2 password req param
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.UserAccount.Password))
	if err != nil {
		err = redis.SetCheckFailCounter(username)
		if err != nil {
			return 0, errorCode.ServiceErr
		}
		return 0, errorCode.LoginErr
	}
	//Login successful Remove failure counter
	err = redis.DeleteLoginFailCounter(username)
	if err != nil {
		return 0, errorCode.ServiceErr
	}

	err = redis.DeleteMessageLatestTime(int64(user.ID))
	if err != nil {
		return 0, errorCode.ServiceErr
	}
	return int64(user.ID), nil
}
