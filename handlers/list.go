package handlers

import (
	"fmt"
	"net/http"

	"github.com/nurali-techie/rc/service"
)

func NewListHandler(service service.CommandService) func(res http.ResponseWriter, req *http.Request) {

	listHandler := func(res http.ResponseWriter, req *http.Request) {
		query := req.URL.Query().Get("query")
		keys, err := service.List(req.Context(), query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, key := range keys {
			key = fmt.Sprintf("%s\n", key)
			res.Write([]byte(key))
		}
	}

	return listHandler
}
