package answers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type QuestionAnswerResponse struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	QuestionID string `json:"question_id"`
}

// GetQuestionAnswerHandler returns an HTTP handler to perform health checks.
func GetQuestionAnswerHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := queryBus.Ask(ctx, fetcher.NewQuestionAnswerQuery(ctx.Param("id")))
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

		answer, ok := resp.(fetcher.AnswerResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, QuestionAnswerResponse{
			ID:         answer.ID(),
			Text:       answer.Text(),
			QuestionID: answer.QuestionID(),
		})
	}
}
