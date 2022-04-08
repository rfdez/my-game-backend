package checking

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/command"
)

const CheckCommandType command.Type = "command.checking.check"

// CheckCommand is a command to check the status of the application.
type CheckCommand struct {
}

// NewCheckCommand creates a new instance of CheckCommand.
func NewCheckCommand() CheckCommand {
	return CheckCommand{}
}

func (c CheckCommand) Type() command.Type {
	return CheckCommandType
}

// CheckCommandHandler is a command handler for CheckCommand.
type CheckCommandHandler struct {
	service Service
}

// NewCheckCommandHandler creates a new instance of CheckCommandHandler.
func NewCheckCommandHandler(service Service) CheckCommandHandler {
	return CheckCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h CheckCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	_, ok := cmd.(CheckCommand)
	if !ok {
		return errors.NewWrongInput("unexpected command")
	}

	return h.service.Status(ctx)
}
