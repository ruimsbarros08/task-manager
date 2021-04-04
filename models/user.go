package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	Roles    []Role `json:"roles" gorm:"many2many:user_role;"`
}
