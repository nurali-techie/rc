package command

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/domain"
)

type moveCommand struct {
	service domain.CommandService
}

func NewMoveCommand(service domain.CommandService) cli.Command {
	moveCmd := new(moveCommand)
	moveCmd.service = service
	return moveCmd
}

func (c *moveCommand) Execute(ctx context.Context, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("parameter missing, please provide <existing_command_name> as first parameter, <new_command_name> as second parameter")
	}

	oldName := args[0]
	newName := args[1]
	err := c.service.Move(ctx, oldName, newName)
	if err != nil {
		return err
	}

	fmt.Println("info: command moved")
	return nil
}
