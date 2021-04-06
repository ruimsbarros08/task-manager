package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ruimsbarros08/task-manager/models"
	"log"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) UserHasRole(user models.User, roleName string) (bool, error) {
	var roles []models.Role
	r.Db.Raw("SELECT * FROM roles r JOIN user_role ur ON ur.role_id = r.id WHERE r.name = ? AND ur.user_id = ?", roleName, user.ID).Scan(&roles)

	log.Printf("Roles for user %d, for role %s. Result %d", user.ID, roleName, len(roles))
	return len(roles) > 0, nil
}
