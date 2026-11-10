package funcs_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/catdevman/gophlings/19_closures"
)

func TestAdder(t *testing.T) {
	a := funcs.Adder()
	b := funcs.Adder()
	if got := a(1); got != 1 {
		t.Fatalf("a(1) = %d, want 1", got)
	}
	if got := b(10); got != 10 {
		t.Fatalf("b(10) = %d, want 10", got)
	}
	if got := a(2); got != 3 {
		t.Fatalf("a(2) = %d, want 3", got)
	}
}

func TestHello(t *testing.T) {
	mux := funcs.Route()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	if rr.Code == http.StatusNotFound {
		t.Fatal(`register Hello with HandleFunc at "/hello"`)
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if got := rr.Body.String(); got != "Hello, gopher" {
		t.Fatalf("body = %q, want %q", got, "Hello, gopher")
	}
}
