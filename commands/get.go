package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/service"
)

type getCommand struct {
	service service.CommandService
	output  cli.Output
}

func NewGetCommand(service service.CommandService, output cli.Output) cli.Command {
	getCmd := new(getCommand)
	getCmd.service = service
	getCmd.output = output
	return getCmd
}

func (c *getCommand) Execute(ctx context.Context, args []string) error {
	fmt.Println("GetCmd, args:", args)
	if len(args) == 0 {
		return nil
	}
	command, err := c.service.Get(ctx, args[0])
	if err != nil {
		return err
	}
	c.output.SetContent([]byte(fmt.Sprintf("key=%s, text=%s", command.Key, command.Text)))
	return nil
}
