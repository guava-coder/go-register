package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "goregister.com/app/auth"
	. "goregister.com/app/controller"
	. "goregister.com/app/db"
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
	NewUserController(user.NewUserService(userRepo, userAuth), router).Run()
	NewJwtController(jwt.NewJwtService(userRepo, userAuth), router).Run()
	NewEmailController(email.NewEmailService(userRepo, sender), router).Run()
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
	port := ":8082"

	err := router.Run(addr + port)
	if err != nil {
		log.Fatal("server start failed" + err.Error())
	}

}

var (
	userRepo user.UserRepository
)

func initRepos() {
	db := DBInit()

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
