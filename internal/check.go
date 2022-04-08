package mygame

import "context"

// CheckRepository defines the expected behaviour from a check repository.
type CheckRepository interface {
	Status(context.Context) error
}
