package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type randomEventResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RandomHandler returns an HTTP handler to perform health checks.
func RandomHandler(queryBus query.Bus[query.Response]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := queryBus.Ask(ctx, fetcher.NewRandomEventQuery("ola"))
		if err != nil {
			if errors.IsWrongInput(err) {
				ctx.AbortWithError(http.StatusBadRequest, errors.WrapWrongInput(err, "invalid date"))
				return
			}
			ctx.AbortWithError(http.StatusInternalServerError, errors.WrapInternal(err, "failed to perform query"))
			return
		}

		event, ok := resp.(fetcher.RandomEventResponse)
		if !ok {
			ctx.AbortWithError(http.StatusInternalServerError, errors.NewInternal("failed to perform query"))
			return
		}

		ctx.JSON(http.StatusOK, randomEventResponse{
			ID:   event.ID(),
			Name: event.Name(),
		})
	}
}
