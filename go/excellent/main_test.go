package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	// Test for even number
	evenVal := 4
	result := EvenOrOdd(evenVal)
	if result != "Even" {
		t.Errorf("Expected Even for %d, actual: %s", evenVal, result)
	}
	// Test for odd 
	oddVal := 5
	result = EvenOrOdd(oddVal)
	if result != "Odd" {
		t.Errorf("Expected Odd for %d, actual: %s", oddVal, result)
	}
}