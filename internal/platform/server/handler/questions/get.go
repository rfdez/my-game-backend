package questions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/kit/query"
)

type questionResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// GetQuestionHandler returns an HTTP handler to perform health checks.
func GetQuestionHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := queryBus.Ask(ctx, fetcher.NewQuestionQuery(ctx.Param("question_id")))
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

		question, ok := resp.(fetcher.QuestionResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		response := questionResponse{
			ID:   question.ID(),
			Text: question.Text(),
		}

		ctx.JSON(http.StatusOK, response)
	}
}
