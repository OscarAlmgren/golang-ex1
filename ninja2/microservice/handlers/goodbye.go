package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye for farewells
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Bye!"))
}
