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
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Bad Request",
				})
				return
			}

			if errors.IsNotFound(err) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Not Found",
				})
				return
			}

			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		event, ok := resp.(fetcher.RandomEventResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
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
