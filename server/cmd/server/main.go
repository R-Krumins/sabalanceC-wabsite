package main

import (
	"fmt"
	"net/http"
	"server/internal/config"
	"server/internal/routers"
)

func main() {
	config := config.LoadConfig()
	mux := http.NewServeMux()

	mux.Handle("/", routers.PageRouter())
	mux.Handle("/api/product", routers.ProductRouter())

	fmt.Printf("Server listening on port %d...\n", config.ServerPort)
	var address = fmt.Sprint(":", config.ServerPort)
	http.ListenAndServe(address, mux)
}
