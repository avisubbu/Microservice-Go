package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Goodbye struct Object
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye Handler constructor
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServerHTTP implements the http.Handler interface
func (h *Goodbye) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Goodbye Request")

	fmt.Fprintf(rw, "Goodbye")
}
