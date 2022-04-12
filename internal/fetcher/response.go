package fetcher

import "github.com/rfdez/my-game-backend/kit/query"

type RandomEventResponse struct {
	id   string
	name string
}

func NewRandomEventResponse(id string, name string) query.Response {
	return RandomEventResponse{
		id:   id,
		name: name,
	}
}

func (r RandomEventResponse) ID() string {
	return r.id
}

func (r RandomEventResponse) Name() string {
	return r.name
}
