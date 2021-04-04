package models

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Summary     string    `json:"summary" gorm:"type:TEXT"`
	PerformedAt time.Time `json:"performed_at"`
	Technician  User      `json:"performed_by"`
}
