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
	res := serv.repo.QueryById(id)
	if res.Id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "User not found",
		})
	} else {
		res.Password = ""
		res.Auth = ""
		ctx.JSON(http.StatusOK, gin.H{
			"Response": "Found User",
			"User":     res,
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
		user := serv.repo.QueryByInfo(us)

		if user.Id == "" {
			psw, err := bcrypt.GenerateFromPassword([]byte(us.Password), 0)
			if err == nil {
				handleAddUser(us, string(psw))
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"Response": "System error, please try again later. ERROR: " + err.Error(),
				})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "User already exist",
			})
		}
	}

	hashUserPasswordAndInsert(user)
}

func (serv UserService) UpdateUserAuth(ctx *gin.Context) {
	changeUserAuthToReal := func(id string) {
		usrWithAuth := User{
			Id:   id,
			Auth: string(serv.UserAuth.MustGetHashAuth()),
		}
		res := serv.repo.UpdateUserAuth(usrWithAuth)
		if res.Auth == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "System failed to generate auth, please try again later",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "User Auth update successful",
			})
		}
	}

	handleUpdate := func(usr User) {
		if serv.repo.CheckAuth(usr) {
			changeUserAuthToReal(usr.Id)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "User ID or Token incorrect",
			})
		}

	}
	serv.readAndHandleRequestBody(ctx, handleUpdate)
}

func (serv UserService) UpdatePassword(ctx *gin.Context) {
	serv.readAndHandleRequestBody(ctx, func(u User) {
		psw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0)
		if err == nil {
			u.Password = string(psw)
			res := serv.repo.UpdatePassword(u)
			if res.Id == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": "User not found",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"Response": "User Password updated",
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "System error. " + err.Error(),
			})
		}

	})
}
