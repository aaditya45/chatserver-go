package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"aadityasharma/golang/chatroom/config"
	 "aadityasharma/golang/chatroom/controller"
	 "aadityasharma/golang/chatroom/middlewares"
	"aadityasharma/golang/chatroom/service"
	"aadityasharma/golang/chatroom/entity"
	"gorm.io/gorm"
	"github.com/gorilla/websocket"
)

var (
	userService    service.UserService       = service.NewUser()
	userController controller.UserController
	db                   *gorm.DB
)

type Session struct {
	ID        string
	user      *entity.User
	ExpiresAt time.Time
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func setupLogOutputInFile() {
	res, _ := os.Create("loggerFile.log")
	gin.DefaultWriter = io.MultiWriter(res, os.Stdout)
}

func init() {
	db = config.GetDBInstance()
	userController = controller.NewUser(userService)
}

func main() {
	go h.run()
	setupLogOutputInFile()
	router := gin.New()

	//cross origin
	corsconfig := cors.DefaultConfig()
	corsconfig.AllowAllOrigins = true
	corsconfig.AllowCredentials = true

	router.LoadHTMLFiles("login.html")
	router.Use(middlewares.Logger(),cors.New(corsconfig))
	
	router.GET("/ws", func(c *gin.Context) {
		serveWs(c.Writer, c.Request, "1")
	})

	router.POST("/register", func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		registeredUser, err := userController.RegisterUser(user.Username, user.Password)
		fmt.Println(registeredUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success":"user registered"})
	})

	router.POST("/login", func(c *gin.Context) {
		var user entity.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		authenticatedUser, err := userController.AuthenticateUser(user.Username, user.Password)
		fmt.Println(authenticatedUser)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		fmt.Println("user logged in")
		c.JSON(http.StatusOK, gin.H{"message":"User logged in","success":true})
	})
	
	router.Run(":8082")
}
