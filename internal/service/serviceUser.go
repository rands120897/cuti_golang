package service

import (
	"cuti/internal/connection"
	"cuti/internal/model"
	"cuti/internal/model/web"
	"cuti/internal/repository"
)

func InputUser(UserRequest web.UserRequest) (model.User, error) {

	userinsert := model.User{
		Name:      UserRequest.Name,
		Email:     UserRequest.Email,
		Password:  UserRequest.Password,
		User_role: UserRequest.User_role,
	}

	_, err := repository.InsertUser(connection.OpenConnection(), userinsert)

	return userinsert, err

}
