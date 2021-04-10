package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/command"
)

type getCmd struct {
}

func NewGetCommand() command.Command {
	return new(getCmd)
}

func (c *getCmd) Execute(ctx context.Context, in command.Input, out command.Output, args []string) error {
	fmt.Println("GetCmd, args:", args)
	return nil
}
