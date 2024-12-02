package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"column:id;type:uint;not null;primaryKey;autoIncrement;comment:主键;"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;type:time;not null;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;type:time;not null;autoUpdateTime;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
