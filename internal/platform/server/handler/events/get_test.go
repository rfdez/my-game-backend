package events_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/events"
	"github.com/rfdez/my-game-backend/kit/query/querymocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Get_ServiceError(t *testing.T) {
	validResponse := fetcher.NewRandomEventResponse(
		"a79f32d8-354f-4063-ad70-a1456fc562bc",
		"test event",
		time.Now().Format(time.RFC3339),
		[]string{"test tag", "test tag 2"},
	)

	queryBus := new(querymocks.Bus)
	queryBus.On(
		"Ask",
		mock.Anything,
		mock.AnythingOfType("fetcher.RandomEventQuery"),
	).Return(validResponse, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/events/random", events.RandomHandler(queryBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/events/random?date=2020-01-01", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
