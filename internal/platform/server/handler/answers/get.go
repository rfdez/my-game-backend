package answers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type eventQuestionAnswerResponse struct {
	EventID    string `json:"event_id"`
	QuestionID string `json:"question_id"`
	Text       string `json:"text"`
}

// GetEventQuestionAnswerHandler returns an HTTP handler to perform health checks.
func GetEventQuestionAnswerHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := queryBus.Ask(ctx, fetcher.NewEventQuestionAnswerQuery(ctx.Param("event_id"), ctx.Param("question_id")))
		if err != nil {
			if errors.IsWrongInput(err) {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			if errors.IsNotFound(err) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		answer, ok := resp.(fetcher.AnswerResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, eventQuestionAnswerResponse{
			EventID:    answer.EventID(),
			QuestionID: answer.QuestionID(),
			Text:       answer.Text(),
		})
	}
}
