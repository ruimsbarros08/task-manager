package mappings

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/services"
	"net/http"
)

type Middleware struct {
	AuthService services.AuthenticationService
	UserService services.UserService
}

func (s *Middleware) AuthMiddleware(c *gin.Context) {
	bearToken := c.Request.Header.Get("Authorization")
	token := s.AuthService.ExtractTokenFromHeader(bearToken)

	userId, err := s.AuthService.ExtractTokenMetadata(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user, err := s.UserService.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

func (s *Middleware) IsManagerMiddleWare(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is missing"})
		return
	}
	user := u.(models.User)

	role, err := s.UserService.UserHasRole(user, "Manager")

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if !role {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User has no permission"})
		c.Abort()
		return
	}

	c.Next()
}

func (s *Middleware) IsTechnicianMiddleware(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is missing"})
		return
	}
	user := u.(models.User)

	role, err := s.UserService.UserHasRole(user, "Technician")

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if !role {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User has no permission"})
		c.Abort()
		return
	}

	c.Next()
}
