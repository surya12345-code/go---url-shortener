package main

import (
	"fmt"
	"net/http"
	"urlshortener/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/shorturl", handlers.ShortURLHandler)
	http.HandleFunc("/bitly/", handlers.RedirectURLHandler)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
