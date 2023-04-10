package common

import "time"

type SQLModel struct {
	Id        uint       `json:"-" gorm:"column:id;primaryKey;autoIncrement:true"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status,omitempty" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
