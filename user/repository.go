package user

import (
	"errors"

	. "goregister.com/app/data"
)

type UserError struct {
	NotFound error
}

func NewUserError() UserError {
	return UserError{
		NotFound: errors.New("user not found"),
	}
}

type UserRepository struct {
	DB map[string]User
}

func NewUserRepository(db map[string]User) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (repo UserRepository) QueryById(id string) (User, error) {
	if repo.DB[id].Id == "" {
		return repo.DB[id], NewUserError().NotFound
	} else {
		return repo.DB[id], nil
	}
}

func (repo UserRepository) QueryByInfo(user User) (User, error) {
	users := make([]User, 0, len(repo.DB))
	for _, v := range repo.DB {
		users = append(users, v)
	}
	for _, v := range users {
		if v.Name == user.Name || v.Email == user.Email {
			return v, nil
		}
	}
	return User{}, NewUserError().NotFound
}

// IsTempCodeCorrect checks if the temporary code of a user is correct.
//
// It takes a User object as a parameter and returns a boolean value indicating
// whether the temporary code is correct. The function queries the user's ID from
// the UserRepository and checks if the temporary code matches the user's temporary
// code in the database. If the temporary code is correct and its length is greater
// than or equal to 6, the function updates the temporary code in the database to an
// empty string and returns true. Otherwise, it returns false.
func (repo UserRepository) IsTempCodeCorrect(user User) bool {
	v, err := repo.QueryById(user.Id)
	if err != nil {
		return false
	}
	if v.TempCode == user.TempCode && len(user.TempCode) >= 6 {
		temp := repo.DB[user.Id]
		temp.TempCode = ""
		repo.DB[user.Id] = temp
		return true
	} else {
		return false
	}
}

func (repo UserRepository) AddUser(user User) User {
	repo.DB[user.Id] = user
	return repo.DB[user.Id]
}

func (repo UserRepository) UpdateUserInfo(user User) (User, error) {
	if repo.DB[user.Id].Id == "" {
		return User{}, NewUserError().NotFound
	} else {
		temp := repo.DB[user.Id]
		temp.Name = user.Name
		temp.Bio = user.Bio

		repo.DB[user.Id] = temp

		return repo.DB[user.Id], nil
	}
}

func (repo UserRepository) UpdatePassword(user User) (User, error) {
	if repo.DB[user.Id].Id == "" {
		return User{}, NewUserError().NotFound
	} else {
		temp := repo.DB[user.Id]
		temp.Password = user.Password
		repo.DB[user.Id] = temp

		return repo.DB[user.Id], nil
	}

}

func (repo UserRepository) DeleteUser(id string) error {
	if repo.DB[id].Id == "" {
		return NewUserError().NotFound
	} else {
		delete(repo.DB, id)
		return nil
	}
}

func (repo UserRepository) UpdateTempCode(user User) (User, error) {
	if user.TempCode == "" {
		return repo.DB[user.Id], errors.New("no tempCode")
	} else {
		temp := repo.DB[user.Id]
		temp.TempCode = user.TempCode
		repo.DB[user.Id] = temp
		return repo.DB[user.Id], nil
	}
}

func (repo UserRepository) UpdateUserAuth(user User) (User, error) {
	if user.Auth == "" {
		return repo.DB[user.Id], errors.New("no auth")
	} else {
		if repo.DB[user.Id].Id == "" {
			return repo.DB[user.Id], NewUserError().NotFound
		} else {
			temp := repo.DB[user.Id]
			temp.Auth = user.Auth
			repo.DB[user.Id] = temp
			return repo.DB[user.Id], nil
		}
	}
}
