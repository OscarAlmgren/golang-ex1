package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handler
type Hello struct {
	l *log.Logger
}

// NewHello with l logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooopla", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Ooops"))
		return
	}

	// log.Printf("Data %s", data)
	fmt.Fprintf(rw, "Hello %s", data)
}
