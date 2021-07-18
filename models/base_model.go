package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id        int64
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// 删除时间可以为NULL
type Model struct {
	BaseModel
	IsDelete  bool         `gorm:"column:is_delete;type:BOOL;notnull"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:DATETIME"`
}
