package services

import (
	"github.com/ruimsbarros08/task-manager/models"
	"time"
)

func CreateUser(email string, name string, password string, roles []uint) (models.User, error) {
	password = hashAndSalt(password)
	user := models.User{Email: email, Name: name, Password: password, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := models.DB.Create(&user).Error

	// TODO verify duplicated role_id's
	for _, roleId := range roles {
		addRole(roleId, &user)
	}

	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := models.DB.Where("email = ?", email).First(&user).Error

	return user, err
}

func GetUserById(id uint) (models.User, error) {
	var user models.User
	err := models.DB.Where("id = ?", id).First(&user).Error

	return user, err
}

func addRole(roleId uint, user* models.User) {
	var role models.Role
	err := models.DB.Where("id = ?", roleId).First(&role).Error

	if err != nil {
		// TODO return error
	}

	models.DB.Model(&user).Association("Roles").Append(&role)
}

