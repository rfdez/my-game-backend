package checks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler returns an HTTP handler to perform health checks.
func PingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	}
}
