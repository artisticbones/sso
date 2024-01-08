package global

import "time"

type Model struct {
	ID        uint `gorm:"primaryKey;type:uint;size:32" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
