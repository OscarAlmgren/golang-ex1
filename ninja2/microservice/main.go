package main

import (
	"log"
	"net/http"
	"os"

	"oscaralmgren.com/oalmgren/microservicetest/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// })

	// http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Goodbye world")
	// })

	http.ListenAndServe(":9090", sm)
}
