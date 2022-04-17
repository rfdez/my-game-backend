package questions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type eventQuestionsByRoundResponse struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	EventID string `json:"event_id"`
}

// GetEventQuestionsByRoundHandler returns an HTTP handler to perform health checks.
func GetEventQuestionsByRoundHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		round, err := strconv.Atoi(ctx.Query("round"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		resp, err := queryBus.Ask(ctx, fetcher.NewEventQuestionsByRoundQuery(ctx.Param("id"), round))
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

		questions, ok := resp.(fetcher.EventQuestionsByRoundResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		response := make([]eventQuestionsByRoundResponse, len(questions.Questions()))
		for i, question := range questions.Questions() {
			response[i] = eventQuestionsByRoundResponse{
				ID:      question.ID(),
				Text:    question.Text(),
				EventID: question.EventID(),
			}
		}

		ctx.JSON(http.StatusOK, response)
	}
}
