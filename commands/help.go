package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
)

type helpCommand struct {
}

func NewHelpCommand() cli.Command {
	return new(helpCommand)
}

func (c *helpCommand) Execute(ctx context.Context, in cli.Input, out cli.Output, args []string) error {
	fmt.Println("rc stands for recall, args:", args)
	return nil
}
