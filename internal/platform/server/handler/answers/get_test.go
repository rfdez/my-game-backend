package answers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/answers"
	"github.com/rfdez/my-game-backend/kit/query/querymocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Get_ServiceError(t *testing.T) {
	validResponse := fetcher.NewAnswerResponse(
		"a79f32d8-354f-4063-ad70-a1456fc562bc",
		"5606b92b-993d-4ab2-8ae9-179afd60ab6a",
		"test event",
	)

	queryBus := new(querymocks.Bus)
	queryBus.On(
		"Ask",
		mock.Anything,
		mock.AnythingOfType("fetcher.EventQuestionAnswerQuery"),
	).Return(validResponse, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/events/:event_id/questions/:question_id/answers", answers.GetEventQuestionAnswerHandler(queryBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events/a79f32d8-354f-4063-ad70-a1456fc562bc/questions/5606b92b-993d-4ab2-8ae9-179afd60ab6a/answers", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("given an invalid url request it returns 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events/questions/answers", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
