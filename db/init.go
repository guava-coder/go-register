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
			Phone:    "0012789908",
			BirthDay: "2000-03-11",
			Gender:   "Male",
			Password: "123",
			Auth:     "$2a$10$GevpiE/I67cDSbfpyRaqv.sEvJa.dYVnnvYjymTMdY2gQ66XLyW.O",
		},
		{
			Id:       "a02",
			Name:     "Lisa",
			Email:    "lisa@mail.com",
			Phone:    "0043289704",
			BirthDay: "1993-02-22",
			Gender:   "Female",
			Password: "123",
			Auth:     "none",
		},
		{
			Id:       "a03",
			Name:     "Max",
			Email:    "max@mail.com",
			Phone:    "006657821",
			BirthDay: "2002-12-03",
			Gender:   "Other",
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
