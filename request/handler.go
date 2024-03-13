package request

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	data "goregister.com/app/data"
)

// Reading the request body and convert to given type for operation.
func ReadAndHandleRequestBody[T data.Data](ctx *gin.Context, operation func(T)) {
	handleBody := func(body []byte, operation func(T), ctx *gin.Context) {
		var us T
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
