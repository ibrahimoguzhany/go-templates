package main

import (
	"fmt"
	"github.com/ibrahimoguzhany/go-templates/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Startting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
