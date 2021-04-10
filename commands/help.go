package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/command"
)

type helpCmd struct {
}

func NewHelpCommand() command.Command {
	return new(helpCmd)
}

func (c *helpCmd) Execute(ctx context.Context, in command.Input, out command.Output, args []string) error {
	fmt.Println("rc stands for recall, args:", args)
	return nil
}
