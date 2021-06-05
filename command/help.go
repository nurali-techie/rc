package command

import (
	"context"
	_ "embed"

	"github.com/nurali-techie/rc/io"
)

//go:embed help.txt
var help string

type helpCommand struct {
	output io.Output
}

func NewHelpCommand(output io.Output) Command {
	helpCmd := new(helpCommand)
	helpCmd.output = output
	return helpCmd
}

func (c *helpCommand) Execute(ctx context.Context, args []string) error {
	c.output.SetContent(help)
	return nil
}
