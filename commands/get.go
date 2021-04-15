package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/service"
)

type getCommand struct {
	service service.CommandService
}

func NewGetCommand(service service.CommandService) cli.Command {
	getCommad := new(getCommand)
	getCommad.service = service
	return getCommad
}

func (c *getCommand) Execute(ctx context.Context, in cli.Input, out cli.Output, args []string) error {
	fmt.Println("GetCmd, args:", args)
	if len(args) == 0 {
		return nil
	}
	command, err := c.service.Get(ctx, args[0])
	if err != nil {
		return err
	}
	out.SetContent([]byte(command.Text))
	return nil
}
