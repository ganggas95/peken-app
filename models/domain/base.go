package domain

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	CreateAt int64          `gorm:"autoCreateTime:milli;" json:"createAt"`
	UpdateAt int64          `gorm:"autoUpdateTime:milli;" json:"updateAt"`
	DeleteAt gorm.DeletedAt `gorm:"index" json:"-"`
}
