package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/data"
)

func DBInit() map[string]User {
	usrs := []User{
		{
			Id:       "a01",
			Name:     "Mark",
			Email:    "mark@mail.com",
			Bio:      "Default User",
			Password: "123",
			Auth:     "$2a$10$GevpiE/I67cDSbfpyRaqv.sEvJa.dYVnnvYjymTMdY2gQ66XLyW.O",
		},
		{
			Id:       "a02",
			Name:     "Lisa",
			Email:    "lisa@mail.com",
			Bio:      "Default User",
			Password: "123",
			Auth:     "none",
		},
		{
			Id:       "a03",
			Name:     "Max",
			Email:    "max@mail.com",
			Bio:      "Default User",
			Password: "123",
			Auth:     "none",
		},
	}

	db := make(map[string]User)
	for _, v := range usrs {
		psw, err := bcrypt.GenerateFromPassword([]byte(v.Password), 0)
		if err != nil {
			log.Println(v.Id + " password hash failed.")
		}
		v.Password = string(psw)
		db[v.Id] = v
	}
	return db
}
