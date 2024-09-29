package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	// Test for even number
	evenVal := 4
	result := EvenOrOdd(evenVal)
	if result != "Even" {
		t.Error("Expected Even for %s, actual: %s", evenVal, result)
	}
	// Test for odd 
	oddVal := 5
	result = EvenOrOdd(oddVal)
	if result != "Odd" {
		t.Error("Expected Odd for %s, actual: %s", oddVal, result)
	}
}