package services

import (
	"net/http"

	"github.com/blac-panda/bkstore-users-api/domain/users"
	"github.com/blac-panda/bkstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
