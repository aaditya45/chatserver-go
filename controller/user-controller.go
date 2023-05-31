package controller

import (
	//"errors"
	"fmt"
	//"rand"
	//"log"
	//"base64"
	//"strconv"
	"strings"
	//"github.com/gin-gonic/gin"
	"aadityasharma/golang/chatroom/entity"
	"aadityasharma/golang/chatroom/service"
)



type UserController interface {
	AuthenticateUser(username, password string) (*entity.User, error)
	RegisterUser(username, password string) (*entity.User, error)
	GenerateSessionID() string
}

type userController struct {
	service service.UserService
}

func NewUser(service service.UserService) UserController{
	return &userController{
		service: service,
	}
}

func (c *userController) AuthenticateUser(username, password string) (*entity.User, error) {
	username=strings.Trim(username, " ")
	
	if len(username)==0 {
		return &entity.User{Username: "error", Password: ""}, nil
	}

	var user entity.User
	user.Username = username
	user.Password = password
	user, err := c.service.AuthenticateUser(user)
	if err != nil {
		return &entity.User{Username: "error", Password:""} , err
	}

	return &user, nil
}

// registerUser registers a new user
func (c *userController) RegisterUser(username, password string) (*entity.User, error) {
	username=strings.Trim(username, " ")
	
	if len(username)==0 {
		return &entity.User{Username: "error", Password: ""}, nil
	}

	var user entity.User
	user.Username = username
	user.Password = password
	user, err := c.service.RegisterUser(user)
	if err != nil {
		return &entity.User{Username: "error", Password:""} , err
	}

	fmt.Println("register")
	return &entity.User{Username: user.Username, Password: user.Password}, nil
}

func (c *userController) GenerateSessionID() string {
	
	return "fffff"
}