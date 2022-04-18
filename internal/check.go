package mygame

import "context"

// CheckRepository defines the expected behaviour from a check repository.
type CheckRepository interface {
	Status(context.Context) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CheckRepository
