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
}

func NewAddCommand(service service.CommandService) cli.Command {
	addCommand := new(addCommand)
	addCommand.service = service
	return addCommand
}

func (c *addCommand) Execute(ctx context.Context, in cli.Input, out cli.Output, args []string) error {
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
