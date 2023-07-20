package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	redisModel "runedance/comment/dao/redis/model"
	"runedance/common/constant"
	"strconv"
)

func getCommentCountById(videoId int64) (int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.CommentCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	result, err := redis.Int64(redisConn.Do("get", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return -1, err
		}
		return 0, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return result, err
	}
	return result, err
}

func GetCommentList(videoId int64) ([]redisModel.CommentRedis, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.CommentRedisPrefix + strconv.FormatInt(videoId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}
	if len(result) <= 0 {
		return nil, redis.ErrNil
	}

	commentList := make([]redisModel.CommentRedis, len(result))
	for i := range result {
		commentInfoKey := constant.CommentInfoRedisPrefix + strconv.FormatInt(int64(result[i]), 10)
		commentInfoResult, commentInfoErr := redis.Bytes(redisConn.Do("get", commentInfoKey))
		if commentInfoErr != nil {
			return nil, commentInfoErr
		}
		commentInfo := new(redisModel.CommentRedis)
		err = json.Unmarshal(commentInfoResult, commentInfo)
		if err != nil {
			return nil, err
		}
		commentList[i] = *commentInfo
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func IsCommentKeyExist(videoId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	//comment:videoId
	key := constant.CommentRedisPrefix + strconv.FormatInt(videoId, 10)
	result, err := redis.Strings(redisConn.Do("keys", key))
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, nil
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return false, err
	}
	return true, nil
}
