package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	. "goregister.com/app/auth"
)

type BearerVerfier struct{}

func isVerified(bearer string, userAuth []byte) bool {
	token := strings.ReplaceAll(bearer, "Bearer ", "")
	var provider JwtProvider
	res := provider.MustIsJWTVerified(userAuth, token)
	return res
}

func handleVerifyFailed(bearer string, ctx *gin.Context) {
	if len(bearer) > 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": "Token expired or not jwt.",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No bearer token found in header.",
		})
	}
}

func handleBearerToken(bearer string, ctx *gin.Context, authOp func(ctx *gin.Context)) {
	var ua UserAuth
	if isVerified(bearer, ua.MustGetOriginAuth()) {
		authOp(ctx)
	} else {
		handleVerifyFailed(bearer, ctx)
	}
}

func (verf BearerVerfier) VerifyBearerToken(ctx *gin.Context, authOp func(ctx *gin.Context)) {
	bearers := ctx.Request.Header["Authorization"]
	if len(bearers) > 0 {
		handleBearerToken(bearers[0], ctx, authOp)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No Authrization",
		})
	}
}

func handleUserIdFromJWT(bearer string, ctx *gin.Context, authOp func(ctx *gin.Context, id string)) {
	var ua UserAuth
	token := strings.ReplaceAll(bearer, "Bearer ", "")
	var provider JwtProvider
	claims := provider.MustGetJWTClaims(ua.MustGetOriginAuth(), token)

	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": "JWT verify failed",
		})
	} else {
		id := fmt.Sprint(claims["id"])
		authOp(ctx, id)
	}
}

func (verf BearerVerfier) ExtractUserIdFromBearer(ctx *gin.Context, authOp func(ctx *gin.Context, id string)) {
	bearers := ctx.Request.Header["Authorization"]
	if len(bearers) > 0 {
		handleUserIdFromJWT(bearers[0], ctx, authOp)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": "No Authrization",
		})
	}
}
