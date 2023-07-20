package entity

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Relation struct {
	ID        uint  `gorm:"primarykey"`
	UserId    int64 `gorm:"column:user_id;not null;index:fk_user_relation"`
	ToUserId  int64 `gorm:"column:to_user_id;not null;index:fk_user_relation_to"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt
}

func (Relation) TableName() string {
	return "relation"
}
