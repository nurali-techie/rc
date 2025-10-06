package command

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/domain"
)

type removeCommand struct {
	service domain.CommandService
}

func NewRemoveCommand(service domain.CommandService) cli.Command {
	removeCmd := new(removeCommand)
	removeCmd.service = service
	return removeCmd
}

func (c *removeCommand) Execute(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("parameter missing, please provide <command_name> as first parameter")
	}

	removeKey := args[0]
	err := c.service.Remove(ctx, removeKey)
	if err != nil {
		return err
	}

	fmt.Println("info: command removed")
	return nil
}
