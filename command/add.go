package command

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/io"
)

type addCommand struct {
	service domain.CommandService
	input   io.Input
}

func NewAddCommand(service domain.CommandService, input io.Input) Command {
	addCmd := new(addCommand)
	addCmd.service = service
	addCmd.input = input
	return addCmd
}

func (c *addCommand) Execute(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("parameter missing, please provide <command_name> as first parameter")
	}
	key := args[0]

	text, err := c.input.GetContent()
	if err != nil {
		return err
	}

	command := &domain.Commmand{
		Key:  key,
		Text: text,
	}
	err = c.service.Add(ctx, command)
	if err != nil {
		return err
	}

	fmt.Println("info: command added")
	return nil
}
