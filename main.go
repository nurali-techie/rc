package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/clipboard"
	"github.com/nurali-techie/rc/commands"
	"github.com/nurali-techie/rc/database"
	"github.com/nurali-techie/rc/service"
	"github.com/nurali-techie/rc/store"
)

func main() {
	var err error

	// setup database
	db, closeFn := database.GetDatabase()
	defer closeFn(db)

	// setup dependency
	commandStore := store.NewCommandStore(db)
	commandService := service.NewCommandService(commandStore)

	// register commands
	helpCmd := commands.NewHelpCommand()
	getCmd := commands.NewGetCommand(commandService)
	commander := cli.NewCommander(helpCmd, getCmd)
	commander.Register("add", commands.NewAddCommand(commandService))
	commander.Register("ls", commands.NewListCommand(commandService))
	commander.Register("web", commands.NewWebCommand(commandService))

	// execute command
	clipboard := clipboard.NewClipboard()
	err = commander.ServeCommand(clipboard, clipboard, os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
