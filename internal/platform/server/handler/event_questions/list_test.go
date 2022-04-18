package eventquestions_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	eventquestions "github.com/rfdez/my-game-backend/internal/platform/server/handler/event_questions"
	"github.com/rfdez/my-game-backend/kit/query/querymocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_List_ServiceError(t *testing.T) {
	validResponse := fetcher.NewEventQuestionResponse(
		"a79f32d8-354f-4063-ad70-a1456fc562bc",
		"db0d8b57-4c4e-493c-9301-05baac8fa65a",
		1,
	)

	queryBus := new(querymocks.Bus)
	queryBus.On(
		"Ask",
		mock.Anything,
		mock.AnythingOfType("fetcher.EventQuestionsByRoundQuery"),
	).Return([]fetcher.EventQuestionResponse{validResponse}, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/events/:event_id/questions", eventquestions.ListEventQuestionsByRoundHandler(queryBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events/a79f32d8-354f-4063-ad70-a1456fc562bc/questions", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("given an invalid url request it returns 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events/questions", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
