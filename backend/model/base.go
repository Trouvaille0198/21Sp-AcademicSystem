package model

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 基础模型 删去了gorm.Model的自增主键
type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
