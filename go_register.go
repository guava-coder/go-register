package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	auth "goregister.com/app/auth"
	controller "goregister.com/app/controller"
	db "goregister.com/app/db"
	email "goregister.com/app/email"
	jwt "goregister.com/app/jwt"
	user "goregister.com/app/user"
)

type GoRegister struct{}

var (
	userAuth = auth.NewUserAuth("./auth.txt")
	sender   = email.NewMailSender("./provider.json")
)

func initControllers(router *gin.Engine) {
	controller.NewUserController(user.NewUserService(userRepo, userAuth), router).Run()
	controller.NewJwtController(jwt.NewJwtService(userRepo, userAuth), router).Run()
	controller.NewEmailController(email.NewEmailService(userRepo, sender), router).Run()
}

func (app GoRegister) Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{"127.0.0.1"})
	initRepos()

	initControllers(router)

	router.StaticFS("./static", http.Dir("static/"))
	index(router)

	addr := "localhost"

	port := getPort()

	err := router.Run(addr + port)
	if err != nil {
		log.Fatal("server start failed" + err.Error())
	}
}

func getPort() string {
	var port string
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	} else {
		port = ":8080"
	}
	return port
}

var (
	userRepo user.UserRepository
)

func initRepos() {
	db := db.DBInit()

	if len(db) > 0 {
		userRepo = user.NewUserRepository(db)
	} else {
		log.Fatal("No DataBase found")
	}
}

func index(r *gin.Engine) {
	path := "./static/*.html"
	r.LoadHTMLGlob(path)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
