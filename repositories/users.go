package repositories

import (
	"github.com/ruimsbarros08/task-manager/models"
	"log"
)

type Result struct {
	count int
}

func UserHasRole(user models.User, roleName string) (bool, error) {
	var result Result
	models.DB.Raw("SELECT COUNT(1) FROM user_role ur JOIN roles r ON ur.role_id = r.id WHERE r.name = ? AND ur.user_id = ?", roleName, user.ID).Scan(&result)

	log.Printf("Count roles for user %d, for role %s. Result %d", user.ID, roleName, result.count)
	return result.count == 1, nil
}
