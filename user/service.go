package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/data"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func readAndHandleRequestBody(ctx *gin.Context, operation func(User)) {
	handleBody := func(body []byte, operation func(User), ctx *gin.Context) {
		var us User
		err := json.Unmarshal(body, &us)
		if err == nil {
			operation(us)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Not a user Error:" + err.Error(),
			})
		}
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		handleBody(body, operation, ctx)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": "Reading request body failed. ERROR: " + err.Error(),
		})
	}
}

func (serv UserService) QueryById(ctx *gin.Context) {
	readAndHandleRequestBody(ctx, func(usr User) {
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

	readAndHandleRequestBody(ctx, hashUserPasswordAndInsert)
}
