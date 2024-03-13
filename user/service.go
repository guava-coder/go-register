package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/auth"
	. "goregister.com/app/data"
	. "goregister.com/app/request"
)

type UserService struct {
	repo UserRepository
	UserAuth
}

func NewUserService(repo UserRepository, userAuth UserAuth) UserService {
	return UserService{
		repo:     repo,
		UserAuth: userAuth,
	}
}

func (serv UserService) readAndHandleRequestBody(ctx *gin.Context, op func(User)) {
	ReadAndHandleRequestBody[User](ctx, op)
}

func (serv UserService) QueryById(ctx *gin.Context, id string) {
	res, err := serv.repo.QueryById(id)
	if err == nil {
		res.Password = ""
		res.Auth = ""
		ctx.JSON(http.StatusOK, gin.H{
			"Response": "Found User",
			"User":     res,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "User not found",
		})
	}
}

func (serv UserService) AddUser(ctx *gin.Context, user User) {
	handleAddUser := func(us User, psw string) {
		us.Password = psw

		uuid := uuid.New()
		us.Id = uuid.String()

		us.Auth = ""

		result := serv.repo.AddUser(us)

		if result.Id == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Reponse": "Add User Failed, system error",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "New User added",
				"User":     result,
			})
		}

	}

	hashUserPasswordAndInsert := func(us User) {
		_, err := serv.repo.QueryByInfo(us)

		if err == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "User already exist",
			})
		} else {
			psw, err := bcrypt.GenerateFromPassword([]byte(us.Password), 0)
			if err == nil {
				handleAddUser(us, string(psw))
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"Response": "System error, please try again later. ERROR: " + err.Error(),
				})
			}
		}
	}

	hashUserPasswordAndInsert(user)
}

func (serv UserService) UpdateUserAuth(ctx *gin.Context) {
	getAuth := func(id string) {
		usrWithAuth := User{
			Id:   id,
			Auth: string(serv.UserAuth.MustGetHashAuth()),
		}
		_, err := serv.repo.UpdateUserAuth(usrWithAuth)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "User Auth update successful",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "System failed to generate auth, please try again later",
			})
		}
	}

	handleUpdate := func(usr User) {
		if serv.repo.IsTempCodeCorrect(usr) {
			getAuth(usr.Id)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "User ID or Token incorrect",
			})
		}

	}
	serv.readAndHandleRequestBody(ctx, handleUpdate)
}

func (serv UserService) UpdatePassword(ctx *gin.Context, id string) {
	handleUpdateRes := func(user User) {
		_, err := serv.repo.UpdatePassword(user)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "User Password updated",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Password update failed",
			})
		}
	}

	hashPassword := func(password string, user User) {
		psw, err := bcrypt.GenerateFromPassword([]byte(password), 0)
		if err == nil {
			user.Password = string(psw)
			handleUpdateRes(user)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "System error. " + err.Error(),
			})
		}
	}

	serv.readAndHandleRequestBody(ctx, func(u User) {
		tempUser, err := serv.repo.QueryById(id)
		if err == nil {
			hashPassword(u.Password, tempUser)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": err.Error(),
			})
		}
	})
}

func (serv UserService) UpdateUserInfo(ctx *gin.Context, id string) {
	serv.readAndHandleRequestBody(ctx, func(u User) {
		u.Id = id
		res, err := serv.repo.UpdateUserInfo(u)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Update successful",
				"User":     res,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Not a user",
			})
		}
	})
}
