package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH MyHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handlerm but it is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH MyHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handlerm but it is %T", v)
	}
}
