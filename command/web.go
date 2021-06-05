package command

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/handler"
)

type webCommand struct {
	service domain.CommandService
}

func NewWebCommand(service domain.CommandService) Command {
	webCmd := new(webCommand)
	webCmd.service = service
	return webCmd
}

func (c *webCommand) Execute(ctx context.Context, args []string) error {
	http.HandleFunc("/", handler.NewHomeHandler())
	http.HandleFunc("/ls", handler.NewListHandler(c.service))
	fmt.Println("open http://localhost:8383")
	return http.ListenAndServe(":8383", nil)
}
