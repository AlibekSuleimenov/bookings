package main

import "testing"

func TestRn(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
