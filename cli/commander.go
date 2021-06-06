package cli

import (
	"context"
)

type Commander struct {
	helpCmd    Command
	defaultCmd Command

	commands map[string]Command
}

func NewCommander(helpCmd Command, defaultCmd Command) *Commander {
	commander := new(Commander)
	commander.defaultCmd = defaultCmd
	commander.helpCmd = helpCmd

	commander.commands = make(map[string]Command)
	commander.commands["help"] = helpCmd

	return commander
}

func (c *Commander) Register(name string, command Command) {
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
