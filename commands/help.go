package commands

import (
	"context"
	"fmt"

	"github.com/nurali-techie/rc/cli"
)

type helpCommand struct {
	output cli.Output
}

func NewHelpCommand(output cli.Output) cli.Command {
	helpCmd := new(helpCommand)
	helpCmd.output = output
	return helpCmd
}

func (c *helpCommand) Execute(ctx context.Context, args []string) error {
	c.output.SetContent([]byte(fmt.Sprintf("rc stands for recall, args:%v", args)))
	return nil
}
