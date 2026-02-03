package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	results := EvenOrOdd(10)
	if results != "Even" {
		t.Errorf("Expected Even but got %s", results)
	}