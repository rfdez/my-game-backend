package fetcher

// RandomEventResponse is the response of a random event.
type RandomEventResponse struct {
	id       string
	name     string
	date     string
	keywords []string
}

// NewRandomEventResponse creates a new instance of RandomEventResponse.
func NewRandomEventResponse(id, name, date string, keywords []string) RandomEventResponse {
	return RandomEventResponse{
		id:       id,
		name:     name,
		date:     date,
		keywords: keywords,
	}
}

// ID returns the id of the random event.
func (r RandomEventResponse) ID() string {
	return r.id
}

// Name returns the name of the random event.
func (r RandomEventResponse) Name() string {
	return r.name
}

// Date returns the date of the random event.
func (r RandomEventResponse) Date() string {
	return r.date
}

// Keywords returns the keywords of the random event.
func (r RandomEventResponse) Keywords() []string {
	return r.keywords
}

// EventQuestionResponse is the response of the event questions by round.
type EventQuestionResponse struct {
	eventID    string
	questionID string
	round      int
}

// NewEventQuestionResponse creates a new instance of EventQuestionResponse.
func NewEventQuestionResponse(eventID, questionID string, round int) EventQuestionResponse {
	return EventQuestionResponse{
		eventID:    eventID,
		questionID: questionID,
		round:      round,
	}
}

// EventID returns the id of the event.
func (eq EventQuestionResponse) EventID() string {
	return eq.eventID
}

// QuestionID returns the id of the question.
func (eq EventQuestionResponse) QuestionID() string {
	return eq.questionID
}

// Round returns the round of the question.
func (eq EventQuestionResponse) Round() int {
	return eq.round
}

// QuestionResponse is the response of a list of questions.
type QuestionResponse struct {
	id   string
	text string
}

// NewQuestionResponse creates a new instance of QuestionResponse.
func NewQuestionResponse(id, text string) QuestionResponse {
	return QuestionResponse{
		id:   id,
		text: text,
	}
}

// ID returns the id of the question.
func (q QuestionResponse) ID() string {
	return q.id
}

// Text returns the text of the question.
func (q QuestionResponse) Text() string {
	return q.text
}

// EventQuestionsByRoundResponse is the response of the questions by round.
type EventQuestionsByRoundResponse struct {
	questions []QuestionResponse
}

// NewEventQuestionsByRoundResponse creates a new instance of EventQuestionsByRoundResponse.
func NewEventQuestionsByRoundResponse(questions []QuestionResponse) EventQuestionsByRoundResponse {
	return EventQuestionsByRoundResponse{
		questions: questions,
	}
}

// Questions returns the list of questions.
func (r EventQuestionsByRoundResponse) Questions() []QuestionResponse {
	return r.questions
}

// AnswerResponse is the response of an answer.
type AnswerResponse struct {
	eventID    string
	questionID string
	text       string
}

// NewAnswerResponse creates a new instance of AnswerResponse.
func NewAnswerResponse(eventID, questionID, text string) AnswerResponse {
	return AnswerResponse{
		eventID:    eventID,
		questionID: questionID,
		text:       text,
	}
}

// EventID returns the id of the answer.
func (a AnswerResponse) EventID() string {
	return a.eventID
}

// QuestionID returns the id of the question.
func (a AnswerResponse) QuestionID() string {
	return a.questionID
}

// Text returns the text of the answer.
func (a AnswerResponse) Text() string {
	return a.text
}
