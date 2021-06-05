package command

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/io"
)

type listCommand struct {
	service domain.CommandService
	output  io.Output
}

func NewListCommand(service domain.CommandService, output io.Output) Command {
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
