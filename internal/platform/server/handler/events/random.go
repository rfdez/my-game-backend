package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type randomEventResponse struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Date     string   `json:"date"`
	Keywords []string `json:"keywords"`
}

// RandomHandler returns an HTTP handler to perform health checks.
func RandomHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := queryBus.Ask(ctx, fetcher.NewRandomEventQuery(ctx.Param("date")))
		if err != nil {
			if errors.IsWrongInput(err) {
				ctx.AbortWithError(http.StatusBadRequest, errors.WrapWrongInput(err, "Bad Request"))
				return
			}

			if errors.IsNotFound(err) {
				ctx.AbortWithError(http.StatusNotFound, errors.WrapNotFound(err, "Not Found"))
				return
			}

			ctx.AbortWithError(http.StatusInternalServerError, errors.WrapInternal(err, "Internal Server Error"))
			return
		}

		event, ok := resp.(fetcher.RandomEventResponse)
		if !ok {
			ctx.AbortWithError(http.StatusInternalServerError, errors.NewInternal("Internal Server Error"))
			return
		}

		ctx.JSON(http.StatusOK, randomEventResponse{
			ID:       event.ID(),
			Name:     event.Name(),
			Date:     event.Date(),
			Keywords: event.Keywords(),
		})
	}
}
