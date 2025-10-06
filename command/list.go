package command

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/io"
)

type listCommand struct {
	service domain.CommandService
	output  io.Output
}

func NewListCommand(service domain.CommandService, output io.Output) cli.Command {
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
	memories, err := c.service.List(ctx, query)
	if err != nil {
		return err
	}

	padding := len(strconv.Itoa(len(memories)))
	seqFormat := fmt.Sprintf("%%%dd", padding)
	for _, memory := range memories {
		c.output.SetContent(fmt.Sprintf(seqFormat+") %s\n", memory.Seq, memory.Key))
	}
	return nil
}
