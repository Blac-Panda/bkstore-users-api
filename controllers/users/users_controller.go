package users

import (
	"net/http"

	"github.com/blac-panda/bkstore-users-api/domain/users"
	"github.com/blac-panda/bkstore-users-api/services"
	"github.com/blac-panda/bkstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

// func FindUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }
