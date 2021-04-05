package tests

import (
	"github.com/ruimsbarros08/task-manager/services"
	"testing"
)

func TestCreateTask(t *testing.T) {
	var roles []uint
	user, uErr := services.CreateUser("test@test.com", "Test", "pwd", roles)
	if uErr!= nil {
		t.Fatalf("User should be created")
	}

	_, tErr := services.CreateTask("summary", user)
	if tErr != nil {
		t.Fatalf("Task should be created")
	}
}
