package handler

import (
	"fmt"
	"net/http"

	"github.com/nurali-techie/rc/domain"
)

func NewListHandler(service domain.CommandService) func(res http.ResponseWriter, req *http.Request) {

	listHandler := func(res http.ResponseWriter, req *http.Request) {
		query := req.URL.Query().Get("query")
		memories, err := service.List(req.Context(), query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, memory := range memories {
			key := fmt.Sprintf("%s\n", memory.Key)
			res.Write([]byte(key))
		}
	}

	return listHandler
}
