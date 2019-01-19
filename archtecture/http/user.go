package http

import (
	"net/http"

	"github.com/takochuu/sandbox/archtecture"
)

type Handlar struct {
	UserService archtecture.UserService
}

func (h *Handlar) ServeHTTP(w http.ResponseWriter, r http.Request) {
	// Do your self
}
