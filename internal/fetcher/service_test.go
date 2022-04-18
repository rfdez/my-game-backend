package fetcher_test

import (
	"context"
	"errors"
	"testing"
	"time"

	mygame "github.com/rfdez/my-game-backend/internal"
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

func Test_FetcherService_RandomEvent_NoEventsFoundError(t *testing.T) {
	var (
		eventRepositoryMock         = new(storagemocks.EventRepository)
		eventQuestionRepositoryMock = new(storagemocks.EventQuestionRepository)
		questionRepositoryMock      = new(storagemocks.QuestionRepository)
		answerRepositoryMock        = new(storagemocks.AnswerRepository)
	)

	eventRepositoryMock.On("SearchAll", mock.Anything).Return([]mygame.Event{}, nil)

	fetcherService := fetcher.NewService(
		eventRepositoryMock,
		eventQuestionRepositoryMock,
		questionRepositoryMock,
		answerRepositoryMock,
	)

	_, err := fetcherService.RandomEvent(context.Background(), "")

	eventQuestionRepositoryMock.AssertExpectations(t)
	assert.EqualError(t, err, "no events found")
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
