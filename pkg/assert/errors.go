package assert

import "testing"

// NoError validates that the provided error is nil and fails the test if the error is nil.
func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}
