package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/service"
)

type addCommand struct {
	service service.CommandService
	input   cli.Input
}

func NewAddCommand(service service.CommandService, input cli.Input) cli.Command {
	addCmd := new(addCommand)
	addCmd.service = service
	addCmd.input = input
	return addCmd
}

func (c *addCommand) Execute(ctx context.Context, args []string) error {
	fmt.Println("AddCmd, args:", args)

	if len(args) < 2 {
		return nil
	}

	command := &domain.Commmand{
		Key:  args[0],
		Text: args[1],
	}
	return c.service.Add(ctx, command)
}
