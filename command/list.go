package command

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/io"
	"github.com/nurali-techie/rc/service"
)

type listCommand struct {
	service service.CommandService
	output  io.Output
}

func NewListCommand(service service.CommandService, output io.Output) cli.Command {
	listCmd := new(listCommand)
	listCmd.service = service
	listCmd.output = output
	return listCmd
}

func (c *listCommand) Execute(ctx context.Context, args []string) error {
	query := ""
	if len(args) > 0 {
		query = args[0]
	}
	keys, err := c.service.List(ctx, query)
	if err != nil {
		return err
	}
	for _, key := range keys {
		c.output.SetContent(fmt.Sprintf("%s\n", key))
	}
	return nil
}
