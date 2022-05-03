package main

import (
	"net/http"
	"web/routes"
)

func main() {
	routes.Rotas()
	http.ListenAndServe(":8000", nil)
}
