package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "goregister.com/app/controller"
	. "goregister.com/app/db"
	jwt "goregister.com/app/jwt"
	user "goregister.com/app/user"
)

type GoRegister struct{}

func (app GoRegister) Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{"127.0.0.1"})
	initRepos()

	NewUserController(user.NewUserService(userRepo), router).Run()
	NewJwtController(jwt.NewJwtService(userRepo), router).Run()

	router.StaticFS("./static", http.Dir("static/"))
	index(router)

	addr := "localhost"
	port := ":8080"

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
