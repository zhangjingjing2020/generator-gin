package models

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"` //创建时间
	UpdatedAt time.Time `json:"updated_at"` //修改时间
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

type OpLogJson struct {
	OperatorDesc    string    `json:"operator_desc"`
	OperatorContent string    `json:"operator_content"`
	OperatorBy      uint      `json:"operator_by,omitempty"`
	StartTime       time.Time `json:"start_time"`
}
