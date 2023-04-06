package common

import "time"

type SQLModel struct {
	Id        uint       `json:"id" gorm:"column:id;unique;not null;"`
	Status    int        `json:"status,omitempty" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
