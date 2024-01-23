package user

import (
	"testing"

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
		user := ur.QueryById("a03")
		if user.Id == "" {
			t.Fatal("User not found.")
		} else {
			t.Log(user)
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
		old := ur.QueryById("a01")
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
		oldCount := len(ur.DB)
		db := ur.DeleteUser(id)
		if oldCount > len(db) {
			t.Log("Delete successful")
		} else {
			t.Fatal()
		}
	})
}

func TestUpdateUserAuth(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		var auth UserAuth
		user := User{
			Id:   "a01",
			Auth: string(auth.MustGetHashAuth()),
		}
		res := ur.UpdateUserAuth(user)
		if res.Auth == "" {
			t.Fatal()
		} else {
			t.Log(res)
		}
	})
}

func TestCheckAuth(t *testing.T) {
	runRepoOperation(func(ur UserRepository) {
		fakeAuth := "ggg"
		user := User{
			Id:   "a02",
			Auth: fakeAuth,
		}
		res := ur.UpdateUserAuth(user)
		if res.Name == "" {
			t.Fatal("user not exist")
		}
		if ur.CheckAuth(User{
			Id:   "a02",
			Auth: "ggg",
		}) {
			t.Log(res)
		} else {
			t.Fatal("Auth incorrect")
		}
	})
}
