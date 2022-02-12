package main

import (
	"fmt"
	"net/http"

	"github.com/ibrahimoguzhany/go-templates/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Startting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
