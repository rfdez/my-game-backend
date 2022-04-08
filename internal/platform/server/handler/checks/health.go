package checks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/checking"
	"github.com/rfdez/my-game-backend/kit/command"
)

// HealthHandler returns an HTTP handler to perform health checks.
func HealthHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := commandBus.Dispatch(ctx, checking.NewCheckCommand())
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
