package commands

import (
	"context"
	_ "embed"

	"github.com/nurali-techie/rc/cli"
)

//go:embed help.txt
var help string

type helpCommand struct {
	output cli.Output
}

func NewHelpCommand(output cli.Output) cli.Command {
	helpCmd := new(helpCommand)
	helpCmd.output = output
	return helpCmd
}

func (c *helpCommand) Execute(ctx context.Context, args []string) error {
	c.output.SetContent(help)
	return nil
}
