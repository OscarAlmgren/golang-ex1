package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", greetHandler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}

func greetHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!")
}
