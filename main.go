package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/clipboard"
	"github.com/nurali-techie/rc/commands"
	"github.com/nurali-techie/rc/console"
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

	// setup input, output
	clipboard := clipboard.NewClipboard()
	console := console.NewConsole()

	// init commander
	helpCmd := commands.NewHelpCommand(console)
	getCmd := commands.NewGetCommand(commandService, clipboard)
	commander := cli.NewCommander(helpCmd, getCmd)

	// register commands
	commander.Register("add", commands.NewAddCommand(commandService, clipboard))
	commander.Register("ls", commands.NewListCommand(commandService, console))
	commander.Register("web", commands.NewWebCommand(commandService))

	// execute command
	err = commander.ServeCommand(os.Args[1:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
