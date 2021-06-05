package main

import (
	"context"

	"github.com/nurali-techie/rc/command"
)

type Commander struct {
	helpCmd    command.Command
	defaultCmd command.Command

	commands map[string]command.Command
}

func NewCommander(helpCmd command.Command, defaultCmd command.Command) *Commander {
	commander := new(Commander)
	commander.defaultCmd = defaultCmd
	commander.helpCmd = helpCmd

	commander.commands = make(map[string]command.Command)
	commander.commands["help"] = helpCmd

	return commander
}

func (c *Commander) Register(name string, command command.Command) {
	c.commands[name] = command
}

func (c *Commander) ServeCommand(args []string) error {
	if len(args) == 0 {
		return c.helpCmd.Execute(context.Background(), args)
	}

	cmdName := args[0]
	cmd := c.commands[cmdName]
	if cmd == nil {
		return c.defaultCmd.Execute(context.Background(), args)
	}

	return cmd.Execute(context.Background(), args[1:])
}
