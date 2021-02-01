package main

import (
	"fmt"
	"net/http"

	"./router"
)

func getHuh(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "huh\n")
}

func main() {
	router.GetRoutes()

	http.ListenAndServe(":8080", nil)
}
