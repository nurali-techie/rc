package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nurali-techie/rc/cfg"
	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/command"
	"github.com/nurali-techie/rc/db"
	"github.com/nurali-techie/rc/io"
	"github.com/nurali-techie/rc/service"
	"github.com/nurali-techie/rc/store"
)

func main() {
	var err error

	// init config
	cfg.Init()

	// setup database
	db, closeFn := db.GetDatabase()
	defer closeFn(db)

	// setup dependency
	commandStore := store.NewCommandStore(db)
	commandService := service.NewCommandService(commandStore)

	// setup input, output
	clipboard := io.NewClipboard()
	console := io.NewConsole()

	// init commander
	helpCmd := command.NewHelpCommand(console)
	getCmd := command.NewGetCommand(commandService, clipboard)
	commander := cli.NewCommander(helpCmd, getCmd)

	// register commands
	commander.Register("add", command.NewAddCommand(commandService, clipboard))
	commander.Register("ls", command.NewListCommand(commandService, console))
	commander.Register("web", command.NewWebCommand(commandService))

	// execute command
	err = commander.ServeCommand(os.Args[1:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
