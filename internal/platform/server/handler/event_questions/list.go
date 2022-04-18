package eventquestions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type eventQuestionByRoundResponse struct {
	EventID    string `json:"event_id"`
	QuestionID string `json:"question_id"`
	Round      int    `json:"round"`
}

// ListEventQuestionsByRoundHandler returns an HTTP handler to perform health checks.
func ListEventQuestionsByRoundHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		round, err := strconv.Atoi(ctx.Query("round"))
		if err != nil {
			round = 0
		}

		resp, err := queryBus.Ask(ctx, fetcher.NewEventQuestionsByRoundQuery(ctx.Param("event_id"), round))
		if err != nil {
			if errors.IsWrongInput(err) {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		eventQuestions, ok := resp.([]fetcher.EventQuestionResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		response := make([]eventQuestionByRoundResponse, len(eventQuestions))
		for i, question := range eventQuestions {
			response[i] = eventQuestionByRoundResponse{
				EventID:    question.EventID(),
				QuestionID: question.QuestionID(),
				Round:      question.Round(),
			}
		}

		ctx.JSON(http.StatusOK, response)
	}
}
