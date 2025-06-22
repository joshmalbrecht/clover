package assert

import "testing"

// Equals compares the expected value with the actual value and fails the test if they are not equal.
func Equals[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("Expected '%v', but got '%v'", expected, actual)
	}
}
