package main

import (
	"fmt"
	"log"
	"net/http"
	"simple_web_server_example/handler"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/hello", handler.HelloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
