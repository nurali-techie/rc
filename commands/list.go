package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/service"
)

type listCommand struct {
	service service.CommandService
}

func NewListCommand(service service.CommandService) cli.Command {
	listCommand := new(listCommand)
	listCommand.service = service
	return listCommand
}

func (c *listCommand) Execute(ctx context.Context, in cli.Input, out cli.Output, args []string) error {
	query := ""
	if len(args) > 0 {
		query = args[0]
	}
	keys, err := c.service.List(ctx, query)
	if err != nil {
		return err
	}
	for _, key := range keys {
		fmt.Println("VN: key=", key)
	}
	return nil
}
