package main

import (
	"net/http"

	"github.com/leonardops91/projetospessoais/loja/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
