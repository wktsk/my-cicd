package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	// Test for even number
	result := EvenOrOdd(4)
	if result != "Even" {
		t.Error("Expected Even for 4, actual: %s", result)
	}
	// Test for odd number
	result = EvenOrOdd(5)
	if result != "Odd" {
		t.Error("Expected Odd for 5, actual: %s", result)
	}
}