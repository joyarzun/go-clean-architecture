package main

import (
	"io"
	"log"
	"net/http"

	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
    log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}