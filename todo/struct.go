package todo

import (
	"time"
)

type ToDoModel struct {
	Id        uint32    `gorm:"primaryKey" json:"id"`
	Task      string    `gorm:"not null" json:"task"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}

type Tabler interface {
	TableName() string
}

func (ToDoModel) TableName() string {
	return "todos"
}

// ======

type ToDoRequest struct {
	Task string `json:"task" validate:"required"`
}

// ======

type BasicResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
