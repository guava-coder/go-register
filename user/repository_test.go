package user

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/auth"
	. "goregister.com/app/data"
)

var repo = NewUserRepository(map[string]User{
	"a01": {
		Id:       "a01",
		Name:     "Mark",
		Email:    "mark@mail.com",
		Bio:      "Default User",
		Password: "123",
		Auth:     "$2a$10$GevpiE/I67cDSbfpyRaqv.sEvJa.dYVnnvYjymTMdY2gQ66XLyW.O",
		TempCode: "none",
	},
	"a02": {
		Id:       "a02",
		Name:     "Lisa",
		Email:    "lisa@mail.com",
		Bio:      "Default User",
		Password: "123",
		Auth:     "none",
		TempCode: "none",
	},
	"a03": {
		Id:       "a03",
		Name:     "John",
		Email:    "john@mail.com",
		Bio:      "Default User",
		Password: "123",
		Auth:     "none",
		TempCode: "none",
	},
})

func TestQueryRepository(t *testing.T) {
	t.Run("test query by id", func(t *testing.T) {
		_, err := repo.QueryById("a03")
		if err != nil {
			t.Fatal("User not found.")
		}
	})
	t.Run("test query by info", func(t *testing.T) {
		_, err := repo.QueryByInfo(User{
			Email: "mark@mail.com",
		})
		if err != nil {
			t.Fatal(err)
		}
	})

}

func TestModifieRepository(t *testing.T) {
	t.Run("test add user", func(t *testing.T) {
		user := User{
			Id:       "a99",
			Name:     "Dora",
			Email:    "dora@mail.com",
			Bio:      "Test User",
			Password: "1234",
			Auth:     "none",
		}
		res := repo.AddUser(user)
		if res.Id == "" {
			t.Fatal("Add user failed")
		} else {
			t.Logf("user count: %d", len(repo.DB))
			t.Log("New user : ", res)
		}
	})
	t.Run("test update user", func(t *testing.T) {
		old, err := repo.QueryById("a01")
		if err != nil {
			t.Fatal(err)
		}
		user := old
		user.Bio = "Fake User"
		user.Auth = "6666"
		_, err = repo.UpdateUserInfo(user)

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test delete user", func(t *testing.T) {
		id := "a01"
		err := repo.DeleteUser(id)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestUserCredentialRepository(t *testing.T) {
	t.Run("test update user auth", func(t *testing.T) {
		auth := NewUserAuth("../auth.txt")
		user := User{
			Id:   "a01",
			Auth: string(auth.MustGetHashAuth()),
		}
		res, err := repo.UpdateUserAuth(user)
		if err != nil {
			t.Fatal(err)
		}
		if res.Auth == "" {
			t.Fatal("Auth incorrect")
		} else {
			t.Log(res)
		}
	})
	t.Run("test is temp code correct", func(t *testing.T) {
		code := "gggggg"
		user := User{
			Id:       "a01",
			TempCode: code,
		}
		res, err := repo.UpdateTempCode(user)
		if err == nil {
			t.Log(res)
			if !repo.IsTempCodeCorrect(user) {
				t.Fatal("temp code incorrect")
			}
		} else {
			t.Fatal(err)
		}
	})
	t.Run("test update password", func(t *testing.T) {
		user := User{
			Id:       "a01",
			Password: "asd123",
		}
		_, err := repo.QueryById(user.Id)
		if err != nil {
			t.Fatal(err)
		}
		psw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
		if err != nil {
			t.Fatal(err)
		}
		user.Password = string(psw)
		_, err = repo.UpdatePassword(user)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test update temp code", func(t *testing.T) {
		user := User{
			Id:       "a02",
			TempCode: "TEST123",
		}
		_, err := repo.UpdateTempCode(user)
		if err != nil {
			t.Fatal(err)
		}
	})
}
