package router

import (
	"net/http"

	"../api"
)

// GetRoutes fuck off
func GetRoutes() {
	http.HandleFunc("/", api.Resolver)
}
