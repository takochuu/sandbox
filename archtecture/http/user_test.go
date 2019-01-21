package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/takochuu/sandbox/archtecture"

	"github.com/takochuu/sandbox/archtecture/mock"
)

func TestHandlar(t *testing.T) {
	var us mock.UserService
	var h Handlar
	h.UserService = &us

	us.UserFn = func(id int) (*archtecture.User, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &archtecture.User{ID: 100, Name: "susy"}, nil
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/100", nil)
	h.ServeHTTP(w, *r)

	if !us.UserInvoked {
		t.Fatal("expected User() to be invoked")
	}
}
