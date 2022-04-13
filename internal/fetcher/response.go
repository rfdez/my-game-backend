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
