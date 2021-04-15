package commands

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nurali-techie/rc/cli"
	"github.com/nurali-techie/rc/handlers"
	"github.com/nurali-techie/rc/service"
)

type webCommand struct {
	service service.CommandService
}

func NewWebCommand(service service.CommandService) cli.Command {
	webCmd := new(webCommand)
	webCmd.service = service
	return webCmd
}

func (c *webCommand) Execute(ctx context.Context, args []string) error {
	http.HandleFunc("/ls", handlers.NewListHandler(c.service))
	fmt.Println("VN: open http://localhost:8383")
	return http.ListenAndServe(":8383", nil)
}
