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

// QuestionResponse is the response of a list of questions.
type QuestionResponse struct {
	id      string
	text    string
	eventID string
}

// NewQuestionResponse creates a new instance of QuestionResponse.
func NewQuestionResponse(id, text, eventID string) QuestionResponse {
	return QuestionResponse{
		id:      id,
		text:    text,
		eventID: eventID,
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

// EvenID returns the id of the event.
func (q QuestionResponse) EventID() string {
	return q.eventID
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
