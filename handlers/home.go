package handlers

import "net/http"

func NewHomeHandler() func(res http.ResponseWriter, req *http.Request) {

	homeHandler := func(res http.ResponseWriter, req *http.Request) {
		content := `
<body>
	<a href="/ls">ls</a>
</body>`
		res.Write([]byte(content))
	}

	return homeHandler
}
