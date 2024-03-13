package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	auth "goregister.com/app/auth"
)

type BearerVerfier struct {
	userAuth auth.UserAuth
	Ctx      *gin.Context
}

func NewBearerVerfier(userAuth auth.UserAuth, ctx *gin.Context) BearerVerfier {
	return BearerVerfier{
		userAuth: userAuth,
		Ctx:      ctx,
	}
}

func isVerified(bearer string, userAuth []byte) bool {
	token := strings.ReplaceAll(bearer, "Bearer ", "")
	var provider JwtProvider
	res := provider.MustIsJWTVerified(userAuth, token)
	return res
}

func (verf BearerVerfier) handleVerifyFailed(bearer string) {
	if len(bearer) > 0 {
		verf.Ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": "Token expired or not jwt.",
		})
	} else {
		verf.Ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No bearer token found in header.",
		})
	}
}

func (verf BearerVerfier) handleBearerToken(bearer string, authOp func(ctx *gin.Context)) {
	if isVerified(bearer, verf.userAuth.MustGetOriginAuth()) {
		authOp(verf.Ctx)
	} else {
		verf.handleVerifyFailed(bearer)
	}
}

func (verf BearerVerfier) VerifyBearerToken(authOp func(ctx *gin.Context)) {
	bearers := verf.Ctx.Request.Header["Authorization"]
	if len(bearers) > 0 {
		verf.handleBearerToken(bearers[0], authOp)
	} else {
		verf.Ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No Authrization",
		})
	}
}

func (verf BearerVerfier) handleUserIdFromJWT(bearer string, authOp func(ctx *gin.Context, id string)) {
	token := strings.ReplaceAll(bearer, "Bearer ", "")
	var provider JwtProvider
	claims := provider.MustGetJWTClaims(verf.userAuth.MustGetOriginAuth(), token)

	if claims["id"] == nil {
		verf.Ctx.JSON(http.StatusForbidden, gin.H{
			"Response": "JWT verify failed",
		})
	} else {
		id := fmt.Sprint(claims["id"])
		authOp(verf.Ctx, id)
	}
}

func (verf BearerVerfier) ExtractUserIdFromBearer(authOp func(ctx *gin.Context, id string)) {
	bearers := verf.Ctx.Request.Header["Authorization"]
	if len(bearers) > 0 {
		verf.handleUserIdFromJWT(bearers[0], authOp)
	} else {
		verf.Ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No Authrization",
		})
	}
}
