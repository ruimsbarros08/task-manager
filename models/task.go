package models

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Summary     string    `json:"summary"`
	PerformedAt time.Time `json:"performed_at"`
	Technician  User      `json:"performed_by"`
}
