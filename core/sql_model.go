package core

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
