package services

import (
	"github.com/jinzhu/gorm"
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/repositories"
	"time"
)

type UserService struct {
	UsersRepository repositories.UserRepository
	Db *gorm.DB
	EncryptionService EncryptionService
}

func (s * UserService) CreateUser(email string, name string, password string, roles []uint) (models.User, error) {
	password = s.EncryptionService.hashAndSalt(password)
	user := models.User{Email: email, Name: name, Password: password, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := s.Db.Create(&user).Error

	// TODO verify duplicated role_id's
	for _, roleId := range roles {
		s.addRole(roleId, &user)
	}

	return user, err
}

func (s * UserService) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := s.Db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (s * UserService)  GetUserById(id uint) (models.User, error) {
	var user models.User
	err := s.Db.Where("id = ?", id).First(&user).Error

	return user, err
}

func (s * UserService) addRole(roleId uint, user* models.User) {
	var role models.Role
	err := s.Db.Where("id = ?", roleId).First(&role).Error

	if err != nil {
		// TODO return error
	}

	s.Db.Model(&user).Association("Roles").Append(&role)
}

func (s * UserService) UserHasRole(user models.User, roleName string) (bool, error) {
	return s.UsersRepository.UserHasRole(user, roleName)
}

