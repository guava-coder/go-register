package user

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/auth"
	. "goregister.com/app/data"
	. "goregister.com/app/db"
)

func runRepoOperation(operation func(UserRepository)) {
	repo := NewUserRepository(DBInit())

	operation(repo)
}

func TestQueryById(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		user, err := ur.QueryById("a03")
		if err == nil {
			t.Log(user)
		} else {
			t.Fatal("User not found.")
		}
	})
}

func TestQueryByInfo(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		res, err := ur.QueryByInfo(User{
			Email: "mark@mail.com",
		})
		if err == nil {
			t.Log(res)
		} else {
			t.Fatal(err)
		}
	})
}

func TestAddUser(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		user := User{
			Id:       "a99",
			Name:     "Dora",
			Email:    "dora@mail.com",
			Bio:      "Test User",
			Password: "1234",
			Auth:     "none",
		}
		res := ur.AddUser(user)
		if res.Id == "" {
			t.Fatal("Add user failed")
		} else {
			t.Logf("user count: %d", len(ur.DB))
			t.Log("New user : ", res)
		}
	})
}

func TestUpdateUserInfo(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		old, err := ur.QueryById("a01")
		if err != nil {
			t.Fatal(err)
		}
		user := old
		user.Bio = "Fake User"
		user.Auth = "6666"
		res := ur.UpdateUserInfo(user)

		if res.Id == "" {
			t.Fatal("Update failed")
		} else {
			t.Log("old user: ", old)
			t.Log("new user: ", res)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		id := "a01"
		err := ur.DeleteUser(id)
		if err == nil {
			t.Log("Delete successful")
		} else {
			t.Fatal(err)
		}
	})
}

func TestUpdateUserAuth(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		auth := NewUserAuth("../auth.txt")
		user := User{
			Id:   "a01",
			Auth: string(auth.MustGetHashAuth()),
		}
		res, err := ur.UpdateUserAuth(user)
		if err != nil {
			t.Fatal(err)
		}
		if res.Auth == "" {
			t.Fatal()
		} else {
			t.Log(res)
		}
	})
}

func TestIsTempCodeCorrect(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		code := "ggg"
		user := User{
			Id:       "a02",
			TempCode: code,
		}
		res, err := ur.UpdateTempCode(user)
		if err == nil {
			if ur.isTempCodeCorrect(User{
				Id:       "a02",
				TempCode: "ggg",
			}) {
				t.Log(res)
			} else {
				t.Fatal("Auth incorrect")
			}
		} else {
			t.Fatal(err)
		}

	})
}

func TestUpdatePassword(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		user := User{
			Id:       "a01",
			Password: "asd123",
		}
		old, err := ur.QueryById(user.Id)
		if err != nil {
			t.Fatal(err)
		}
		psw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
		if err != nil {
			t.Fatal(err)
		}
		user.Password = string(psw)
		res := ur.UpdatePassword(user)
		if res.Id == "" {
			t.Fatal()
		} else {
			t.Log("old: ", old)
			t.Log("new: ", res)
		}
	})
}

func TestUpdateTempCode(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		user := User{
			Id:       "a02",
			TempCode: "TEST123",
		}
		res, err := ur.UpdateTempCode(user)
		if err == nil {
			t.Log(res)
		} else {
			t.Fatal()
		}
	})
}
