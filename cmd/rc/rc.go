package main

import (
	"fmt"
	"os"

	"github.com/nurali-techie/rc/clipboard"
	"github.com/nurali-techie/rc/command"
	"github.com/nurali-techie/rc/commands"
)

func main() {
	// register commands
	helpCmd := commands.NewHelpCommand()
	getCmd := commands.NewGetCommand()
	commander := command.NewCommander(helpCmd, getCmd)

	// execute command
	clipboard := clipboard.NewClipboard()
	err := commander.ServeCommand(clipboard, clipboard, os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
