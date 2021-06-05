package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nurali-techie/rc/command"
	"github.com/nurali-techie/rc/config"
	"github.com/nurali-techie/rc/database"
	"github.com/nurali-techie/rc/io"
	"github.com/nurali-techie/rc/repository"
	"github.com/nurali-techie/rc/service"
)

func main() {
	var err error

	// init config
	config.Init()

	// setup database
	db, closeFn := database.GetDatabase()
	defer closeFn(db)

	// setup dependency
	commandRepository := repository.NewCommandRepository(db)
	commandService := service.NewCommandService(commandRepository)

	// setup input, output
	clipboard := io.NewClipboard()
	console := io.NewConsole()

	// init commander
	helpCmd := command.NewHelpCommand(console)
	getCmd := command.NewGetCommand(commandService, clipboard)
	commander := NewCommander(helpCmd, getCmd)

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
