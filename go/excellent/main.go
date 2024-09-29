package main

// EvenOrOdd returns "Even" if the given number is even, and "Odd" if the number is odd.
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}