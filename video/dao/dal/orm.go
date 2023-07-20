package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"math"
	"runedance/video/dao/dal/entity"
	"time"
)

// CreateVideo 创建视频
func CreateVideo(ctx context.Context, userId int64, title string, playUrl string, coverUrl string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "dalCreateVideo")
	defer span.Finish()
	video := &entity.Video{
		UserId:   userId,
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&video).Error // 通过数据的指针来创建，所以要用&comment
		if err != nil {
			klog.Error("create comment fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", userId).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
		if err != nil {
			klog.Error("Add user work count error " + err.Error())
			return err
		}
		return nil
	})
	return err
}

// MGetVideoByUserID 根据用户id查视频
func MGetVideoByUserID(ctx context.Context, userId int64) ([]*entity.Video, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "dalMGetVideoByUserID")
	defer span.Finish()
	var videos []*entity.Video
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// GetLikeCount 返回视频点赞数
func GetLikeCount(ctx context.Context, videoID int64) (int64, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "dalGetLikeCount")
	defer span.Finish()
	var video entity.Video
	if err := DB.WithContext(ctx).Where("ID = ?", videoID).First(&video).Error; err != nil {
		return 0, err
	}
	return video.FavoriteCount, nil
}

// GetCommentCount 返回视频评论数
func GetCommentCount(ctx context.Context, videoID int64) (int64, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "dalGetCommentCount")
	defer span.Finish()
	var video entity.Video
	if err := DB.WithContext(ctx).Where("ID = ?", videoID).First(&video).Error; err != nil {
		return 0, err
	}
	return video.CommentCount, nil
}

// IsFavorite 返回是否点赞
func IsFavorite(ctx context.Context, videoID int64, userId int64) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "dalIsFavorite")
	defer span.Finish()
	var favorites []*entity.Favorite
	result := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoID).Find(&favorites)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// MGetVideoByTime 根据时间戳返回最近count个视频,还需要返回next time
func MGetVideoByTime(ctx context.Context, latestTime time.Time, count int64) ([]*entity.Video, int64, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "dalGetVideoByTime")
	defer span.Finish()
	var videos []*entity.Video
	if err := DB.WithContext(ctx).Where("created_at < ?", latestTime).Limit(int(count)).Order("created_at DESC").Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	var nextTime int64 = math.MaxInt32
	if len(videos) != 0 { // 查到了新视频
		nextTime = videos[0].CreatedAt.Unix()
	}
	return videos, nextTime, nil
}

func LikeVideo(ctx context.Context, userID int64, videoID int64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "LikeVideo")
	defer span.Finish()
	favorite := &entity.Favorite{
		UserId:  userID,
		VideoId: videoID,
	}

	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&favorite).Error // 通过数据的指针来创建，所以要用&comment
		if err != nil {
			klog.Error("create favorite fail " + err.Error())
			return err
		}
		// 这里需要指定Table("video")，因为没有model，无法自动确认表名
		err = tx.Table("video").Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			klog.Error("Add video favorite count error " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			klog.Error("Add user favorite count error " + err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UnLikeVideo 取消点赞视频
func UnLikeVideo(ctx context.Context, userID int64, videoID int64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "UnLikeVideo")
	defer span.Finish()
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("favorite").Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&entity.Favorite{}).Error
		if err != nil {
			klog.Error("delete favorite fail: " + err.Error())
			return err
		}
		// 这里需要指定Table("video")，因为没有model，无法自动确认表名
		err = tx.Table("video").Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			klog.Error("Sub video favorite count error " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			klog.Error("Sub user favorite count error " + err.Error())
			return err
		}
		return nil
	})
	return err
}

// MGetLikeList 通过用户ID获取用户点赞的视频ID数组
func MGetLikeList(ctx context.Context, userId int64) ([]int64, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "MGetLikeList")
	defer span.Finish()
	var favorites []*entity.Favorite
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&favorites).Error; err != nil {
		return nil, err
	}
	var likeList []int64
	for _, favorite := range favorites {
		likeList = append(likeList, favorite.VideoId)
	}
	return likeList, nil
}

// MGetVideoInfo 通过视频ID查询得到model.Video信息
func MGetVideoInfo(ctx context.Context, videoID int64) (*entity.Video, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "MGetVideoInfo")
	defer span.Finish()
	var videoInfo *entity.Video
	DB.WithContext(ctx).Where("ID = ?", videoID).First(&videoInfo)
	return videoInfo, nil
}
