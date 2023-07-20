package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"runedance/comment/dao/dal/entity"
	"runedance/kitexGen/kitex_gen/commentproto"
)

func CreateComment(ctx context.Context, userId int64, videoId int64, content string, commentUUID int64, createTime int64) (*commentproto.CommentInfo, error) {
	comment := entity.Comment{
		UserId:      userId,
		VideoId:     videoId,
		Contents:    content,
		CommentUUID: commentUUID,
		CreateTime:  createTime,
	}
	// 创建评论 和 comment_count+1 要在一个Transaction事务中完成
	// 且使用事务的返回值
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&comment).Error // 通过数据的指针来创建，所以要用&comment
		if err != nil {
			klog.Error("create comment fail " + err.Error())
			return err
		}
		// 这里需要指定Table("video")，因为没有model，无法自动确认表名
		err = tx.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			klog.Error("AddCommentCount error " + err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pack.Comment(&comment), nil
}

// DeleteComment
// delete a comment by commentID
func DeleteComment(ctx context.Context, commentID int64, videoID int64) error {
	// 删除评论 和 comment_count-1 要在一个Transaction事务中完成
	// 且使用事务的返回值
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("comment_uuid = ?", commentID).Delete(&entity.Comment{}).Error
		// UPDATE `comment` SET `deleted_at`='\now' WHERE comment_uuid = commentID AND `comment`.`deleted_at` IS NULL
		if err != nil {
			klog.Error("delete comment fail: " + err.Error())
			return err
		}
		// 这里需要指定Table("video")，因为没有model，无法自动确认表名
		err = tx.Table("video").Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			klog.Error("SubCommentCount error " + err.Error())
			return err
		}
		return nil
	})
	return err
}

// GetCommentList
// get comment list by videoID
func GetCommentList(ctx context.Context, videoID int64) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	// 按照评论发布时间降序排序，使用Order("created_at desc")
	err := DB.WithContext(ctx).Where("video_id = ?", videoID).Order("created_at desc").Find(&comments).Error
	if err != nil {
		klog.Error("get comment list fail: " + err.Error())
		return nil, err
	}
	return comments, nil
}
