package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Summary     string    `json:"summary" gorm:"type:TEXT"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	PerformedAt time.Time `json:"performed_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
}
