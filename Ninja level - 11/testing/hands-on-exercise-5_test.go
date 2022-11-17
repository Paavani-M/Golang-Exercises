package main

import "testing"

func TestSquare(t *testing.T) {
	a := 1
	if a != 2 {
		t.Errorf("Expected 2 got %d", a)
	}
}

// basic example of testing working.
