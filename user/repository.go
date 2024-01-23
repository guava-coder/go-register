package user

import (
	. "goregister.com/app/data"
)

type UserRepository struct {
	DB map[string]User
}

func NewUserRepository(db map[string]User) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (repo UserRepository) QueryById(id string) User {
	return repo.DB[id]
}

func (repo UserRepository) QueryByInfo(user User) User {
	users := make([]User, 0, len(repo.DB))
	for _, v := range repo.DB {
		users = append(users, v)
	}
	var target User
	for _, v := range users {
		if v.Name == user.Name || v.Email == user.Email {
			target = v
			break
		}
	}
	return target
}

func (repo UserRepository) CheckAuth(user User) bool {
	v := repo.QueryById(user.Id)
	if v.Auth == user.Auth {
		return true
	} else {
		return false
	}
}

func (repo UserRepository) AddUser(user User) User {
	repo.DB[user.Id] = user
	return repo.DB[user.Id]
}

func (repo UserRepository) UpdateUserInfo(user User) User {
	temp := repo.DB[user.Id]
	temp.Name = user.Name
	temp.Bio = user.Bio

	repo.DB[user.Id] = temp

	return repo.DB[user.Id]
}

func (repo UserRepository) DeleteUser(id string) map[string]User {
	delete(repo.DB, id)
	return repo.DB
}

func (repo UserRepository) UpdateUserAuth(user User) User {
	temp := repo.DB[user.Id]
	temp.Auth = user.Auth
	repo.DB[user.Id] = temp
	return repo.DB[user.Id]
}
