package entity

import (
	"time"
)

type CanceledUsers struct {
	ID        uint32    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserId    uint32    `gorm:"not null" json:"user_id"`
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp null on update now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp null" json:"deleted_at"`
}