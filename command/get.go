package command

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/io"
)

type getCommand struct {
	service domain.CommandService
	output  io.Output
}

func NewGetCommand(service domain.CommandService, output io.Output) Command {
	getCmd := new(getCommand)
	getCmd.service = service
	getCmd.output = output
	return getCmd
}

func (c *getCommand) Execute(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("parameter missing, please provide <number> as first parameter")
	}

	command, err := c.service.Get(ctx, args[0])
	if err != nil {
		return err
	}

	err = c.output.SetContent(command.Text)
	if err != nil {
		return err
	}

	fmt.Println("info: command copied to clipboard")
	return nil
}
