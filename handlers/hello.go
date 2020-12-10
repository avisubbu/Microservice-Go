package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct object
type Hello struct {
	l *log.Logger
}

// NewHello Constructor to create Hello Handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServerHTTP method implements http.Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello Requests")

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.l.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", b)
}
