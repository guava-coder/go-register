package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/data"
	. "goregister.com/app/request"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (serv UserService) readAndHandleRequestBody(ctx *gin.Context, op func(User)) {
	ReadAndHandleRequestBody[User](ctx, op)
}

func (serv UserService) QueryById(ctx *gin.Context) {
	serv.readAndHandleRequestBody(ctx, func(usr User) {
		res := serv.repo.QueryById(usr.Id)
		if res.Id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "User not found",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Find User",
				"User":     res,
			})
		}
	})
}

func (serv UserService) AddUser(ctx *gin.Context) {
	handleAddUser := func(us User, psw string) {
		us.Password = psw

		uuid := uuid.New()
		us.Id = uuid.String()

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

	serv.readAndHandleRequestBody(ctx, hashUserPasswordAndInsert)
}
