package checks_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/checks"
	"github.com/rfdez/my-game-backend/kit/command/commandmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Health(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On("Dispatch", mock.Anything, mock.AnythingOfType("checking.CheckCommand")).Return(nil)
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", checks.HealthHandler(commandBus))

	t.Run("it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
