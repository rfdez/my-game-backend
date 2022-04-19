package fetcher_test

import (
	"context"
	"testing"
	"time"

	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_FetcherService_RandomEvent_RepositoryError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	eventRepositoryMock.On("SearchAll", mock.Anything).Return(nil, errors.New("error"))

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err := fetcherService.RandomEvent(context.Background(), "")

	eventQuestionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_FetcherService_RandomEvent_InvalidArgumentError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err := fetcherService.RandomEvent(context.Background(), "test")

	assert.EqualError(t, err, "invalid test date, the format should be 2006-01-02")
}

func Test_FetcherService_RandomEvent_NoEventsFoundForDateError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	evt, err := mygame.NewEvent(
		"4093dceb-f34a-42c3-bfa6-9344a5c948a3",
		"test",
		time.Now().Format(mygame.EventDateRFC3339),
		[]string{"test", "test2"},
	)
	require.NoError(t, err)

	repoResponse := []mygame.Event{evt}

	eventRepositoryMock.On("SearchAll", mock.Anything).Return(repoResponse, nil)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err = fetcherService.RandomEvent(context.Background(), "2020-01-01")

	eventQuestionRepositoryMock.AssertExpectations(t)
	assert.EqualError(t, err, "no events found for date 2020-01-01")
}

func Test_FetcherService_RandomEvent_Succeed(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	evt, err := mygame.NewEvent(
		"4093dceb-f34a-42c3-bfa6-9344a5c948a3",
		"test",
		time.Now().Format(mygame.EventDateRFC3339),
		[]string{"test", "test2"},
	)
	require.NoError(t, err)

	repoResponse := []mygame.Event{evt}

	eventRepositoryMock.On("SearchAll", mock.Anything).Return(repoResponse, nil)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err = fetcherService.RandomEvent(context.Background(), time.Now().Format(mygame.EventDateRFC3339))

	eventQuestionRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_FetcherService_EventQuestionsByRound_RepositoryError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	eventQuestionRepositoryMock.On("SearchByEventID", mock.Anything, mock.AnythingOfType("mygame.EventID")).Return(nil, errors.New("error"))

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err := fetcherService.EventQuestionsByRound(context.Background(), "4093dceb-f34a-42c3-bfa6-9344a5c948a3", 1)

	eventQuestionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_FetcherService_EventQuestionsByRound_InvalidArgumentError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	t.Run("invalid event id", func(t *testing.T) {
		_, err := fetcherService.EventQuestionsByRound(context.Background(), "test", 1)

		assert.EqualError(t, err, "invalid event id test")
	})

	t.Run("invalid round", func(t *testing.T) {
		_, err := fetcherService.EventQuestionsByRound(context.Background(), "test", 0)

		assert.EqualError(t, err, "round must be greater than 0")
	})
}

func Test_FetcherService_EventQuestionsByRound_Succeed(t *testing.T) {
	t.Run("with questions", func(t *testing.T) {
		var (
			eventRepositoryMock         = new(storagemocks.EventRepository)
			eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
			questionRepositoryMock      = new(storagemocks.QuestionRepository)
			answerRepositoryMock        = new(storagemocks.AnswerRepository)
		)

		evtq, err := mygame.NewEventQuestion(
			"4093dceb-f34a-42c3-bfa6-9344a5c948a3",
			"c9319cb8-1593-46d4-bf7e-9aa24def556a",
			1,
		)
		require.NoError(t, err)

		repoResponse := []mygame.EventQuestion{evtq}

		eventQuestionRepositoryMock.On("SearchByEventID", mock.Anything, mock.AnythingOfType("mygame.EventID")).Return(repoResponse, nil)

		fetcherService := fetcher.NewService(
			eventRepositoryMock,
			eventQuestionRepositoryMock,
			questionRepositoryMock,
			answerRepositoryMock,
		)

		_, err = fetcherService.EventQuestionsByRound(context.Background(), "4093dceb-f34a-42c3-bfa6-9344a5c948a3", 1)

		eventQuestionRepositoryMock.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("with no questions", func(t *testing.T) {
		var (
			eventRepositoryMock         = new(storagemocks.EventRepository)
			eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
			questionRepositoryMock      = new(storagemocks.QuestionRepository)
			answerRepositoryMock        = new(storagemocks.AnswerRepository)
		)

		eventQuestionRepositoryMock.On("SearchByEventID", mock.Anything, mock.AnythingOfType("mygame.EventID")).Return(nil, nil)

		fetcherService := fetcher.NewService(
			eventRepositoryMock,
			eventQuestionRepositoryMock,
			questionRepositoryMock,
			answerRepositoryMock,
		)

		resp, err := fetcherService.EventQuestionsByRound(context.Background(), "4093dceb-f34a-42c3-bfa6-9344a5c948a3", 1)

		eventQuestionRepositoryMock.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Empty(t, resp)
	})
}

func Test_FetcherService_Question_RepositoryError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	questionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("mygame.QuestionID")).Return(mygame.Question{}, errors.New("error"))

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err := fetcherService.Question(context.Background(), "c9319cb8-1593-46d4-bf7e-9aa24def556a")

	questionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_FetcherService_Question_InvalidArgumentError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	t.Run("invalid event id", func(t *testing.T) {
		_, err := fetcherService.Question(context.Background(), "test")

		assert.EqualError(t, err, "invalid question id test")
	})
}

func Test_FetcherService_Question_Succeed(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	q, err := mygame.NewQuestion(
		"4093dceb-f34a-42c3-bfa6-9344a5c948a3",
		"text 1",
	)
	require.NoError(t, err)

	questionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("mygame.QuestionID")).Return(q, nil)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err = fetcherService.Question(context.Background(), "4093dceb-f34a-42c3-bfa6-9344a5c948a3")

	questionRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_FetcherService_EventQuestionAnswer_RepositoryError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	answerRepositoryMock.On("FindByEventIDAndQuestionID", mock.Anything, mock.AnythingOfType("mygame.EventID"), mock.AnythingOfType("mygame.QuestionID")).Return(mygame.Answer{}, errors.New("error"))

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err := fetcherService.EventQuestionAnswer(context.Background(), "c9319cb8-1593-46d4-bf7e-9aa24def556a", "4093dceb-f34a-42c3-bfa6-9344a5c948a3")

	answerRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_FetcherService_EventQuestionAnswer_InvalidArgumentError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	t.Run("invalid event id", func(t *testing.T) {
		_, err := fetcherService.EventQuestionAnswer(context.Background(), "test", "c9319cb8-1593-46d4-bf7e-9aa24def556a")

		assert.EqualError(t, err, "invalid event id test")
	})

	t.Run("invalid question id", func(t *testing.T) {
		_, err := fetcherService.EventQuestionAnswer(context.Background(), "c9319cb8-1593-46d4-bf7e-9aa24def556a", "test")

		assert.EqualError(t, err, "invalid question id test")
	})
}

func Test_FetcherService_EventQuestionAnswer_Succeed(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	a, err := mygame.NewAnswer(
		"4093dceb-f34a-42c3-bfa6-9344a5c948a3",
		"c9319cb8-1593-46d4-bf7e-9aa24def556a",
		"text 1",
	)
	require.NoError(t, err)

	answerRepositoryMock.On("FindByEventIDAndQuestionID", mock.Anything, mock.AnythingOfType("mygame.EventID"), mock.AnythingOfType("mygame.QuestionID")).Return(a, nil)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err = fetcherService.EventQuestionAnswer(context.Background(), "4093dceb-f34a-42c3-bfa6-9344a5c948a3", "c9319cb8-1593-46d4-bf7e-9aa24def556a")

	answerRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
