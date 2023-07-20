package redis

import (
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomodule/redigo/redis"
	"runedance/common/constant"
	"runedance/video/dao/dal/entity"
	redismodel "runedance/video/dao/redis/model"
	"strconv"
)

func GetPublishList(userId int64) ([]*entity.Video, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.PublishRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		return nil, err
	}

	videoRedisListp := new([]redismodel.VideoRedis)
	err = json.Unmarshal(result, videoRedisListp)
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}

	videoRedisList := *videoRedisListp
	videoList := make([]*entity.Video, len(videoRedisList))

	for i := range videoRedisList {
		videoId := videoRedisList[i].VideoId
		cntKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(int64(videoId), 10)
		cnt, cntErr := redis.Int64s(redisConn.Do("hmget",
			cntKey, constant.LikeCountRedisPrefix, constant.CommentCountRedisPrefix))
		if cntErr != nil {
			return nil, cntErr
		}
		videoList[i] = &entity.Video{
			UserId:        videoRedisList[i].UserId,
			Title:         videoRedisList[i].Title,
			PlayUrl:       videoRedisList[i].PlayUrl,
			CoverUrl:      videoRedisList[i].CoverUrl,
			FavoriteCount: cnt[0],
			CommentCount:  cnt[1],
		}
		videoList[i].ID = videoRedisList[i].VideoId
		redisConn.Do("expire", cntKey, expireTime)
	}

	return videoList, nil
}

func GetVideoInfo(videoId int64) (*entity.Video, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.VideoInfoRedisPrefix + strconv.FormatInt(videoId, 10)
	result, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		return nil, err
	}
	videoInfo := new(entity.Video)
	err = json.Unmarshal(result, videoInfo)
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		klog.Error(err)
	}

	cntKey := constant.VideoInfoCntRedisPrefix + strconv.FormatInt(int64(videoId), 10)
	cnt, cntErr := redis.Int64s(redisConn.Do("hmget",
		cntKey, constant.LikeCountRedisPrefix, constant.CommentCountRedisPrefix))
	if cntErr != nil {
		return nil, cntErr
	}
	if cnt == nil || len(cnt) == 0 {
		return nil, redis.ErrNil
	}
	videoInfo.FavoriteCount = cnt[0]
	videoInfo.CommentCount = cnt[1]
	_, err = redisConn.Do("expire", cntKey, expireTime)
	if err != nil {
		klog.Error(err)
	}

	return videoInfo, nil
}

func GetLikeList(userId int64) ([]int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)
	result, err := redis.Int64s(redisConn.Do("zrevrange", key, 0, -1))
	if err != nil {
		return nil, err
	}

	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetIsLikeById(userId int64, videoId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)
	likeKey := strconv.FormatInt(videoId, 10)
	result, err := redisConn.Do("zscore", key, likeKey)
	if err != nil {
		return false, err
	}
	expireTime := expireTimeUtil.GetRandTime()
	_, err = redisConn.Do("expire", key, expireTime)
	if err != nil {
		return false, err
	}
	// ZSCORE key member
	// If the member element is not a member of the ordered set key, or if the key does not exist, result is nil.
	if result == nil {
		return false, nil
	}
	return true, nil
}

func IsLikeKeyExist(userId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	key := constant.LikeRedisPrefix + strconv.FormatInt(userId, 10)
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

func GetLikeCountById(videoId int64) (int64, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	key := constant.LikeCountRedisPrefix + strconv.FormatInt(int64(videoId), 10)
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

func GetCommentCountById(videoId int64) (int64, error) {
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

func IsKeyExistByBloom(prefix string, keyId int64) (bool, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	return true, nil

	result, err := redis.Int(redisConn.Do("bf.exists", prefix, keyId))
	if err != nil {
		return true, err
	}
	if result == 0 {
		return false, nil
	}
	return true, nil
}
