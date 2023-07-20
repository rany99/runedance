package entity

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Favorite Gorm Data Structures
type Favorite struct {
	ID        uint  `gorm:"primarykey"`
	UserId    int64 `gorm:"column:user_id;not null;index:fk_user_favorite"`
	VideoId   int64 `gorm:"column:video_id;not null;index:fk_video_favorite"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt
}

func (f *Favorite) TableName() string {
	return "favorite"
}
