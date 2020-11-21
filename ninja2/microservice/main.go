package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooopla", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops"))
			return
		}

		// log.Printf("Data %s", data)
		fmt.Fprintf(rw, "Hello %s", data)
	})
	http.ListenAndServe(":8080", nil)
}
