package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "goregister.com/app/controller"
	. "goregister.com/app/db"
	email "goregister.com/app/email"
	jwt "goregister.com/app/jwt"
	user "goregister.com/app/user"
)

type GoRegister struct{}

func initControllers(router *gin.Engine) {
	NewUserController(user.NewUserService(userRepo), router).Run()
	NewJwtController(jwt.NewJwtService(userRepo), router).Run()
	NewEmailController(email.NewEmailService(userRepo), router).Run()
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
