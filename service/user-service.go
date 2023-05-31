package service

import (
	"fmt"

	"aadityasharma/golang/chatroom/config"
	"aadityasharma/golang/chatroom/entity"
	"gorm.io/gorm"
	//"github.com/gin-gonic/gin"
)

var(
	db *gorm.DB
)

func init() {
	db = config.GetDBInstance()
	db.AutoMigrate(&entity.User{})
}

type UserService interface {
	AuthenticateUser(entity.User) (entity.User, error)
	RegisterUser(entity.User) (entity.User, error)
	GenerateSessionID() string
}

type userService struct {
	users []entity.User
}

func NewUser() UserService {
	return &userService{}
}

// AuthenticateUser authenticates a user
func (c *userService) AuthenticateUser(user entity.User) (entity.User, error) {
	fmt.Println(user)
	var user_temp entity.User
	if err := db.Where("username = ?", user.Username).First(&user_temp).Error; err != nil {
		fmt.Println("--------------->ERRROR",err)
		return user, err
	}	
	fmt.Println("---------------->",user_temp)

	if user_temp.Password != user.Password {
		return user, fmt.Errorf("wrong credentials")
	}
	return user_temp, nil
}

// registerUser registers a new user
func (c *userService) RegisterUser(user entity.User) (entity.User, error) {
	var user_temp entity.User
	if err := db.Where("username = ?", user.Username).Find(&user_temp); err != nil {
		if err.Error == gorm.ErrRecordNotFound {
			c.users=append(c.users, user)
			db.Create(&user)
			return user,nil
		}
	}
	
	
	return user, fmt.Errorf("user already exists")
}

func (c *userService) GenerateSessionID() string {
	
	return "fffff"
}