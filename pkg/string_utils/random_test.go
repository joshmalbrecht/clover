package string_utils

import (
	"slices"
	"testing"

	"github.com/joshmalbrecht/clover/pkg/assert"
)

type testCharSet struct{}

func (t testCharSet) Characters() []rune {
	return []rune{'a', 'b', 'c'}
}

type testEmptyCharSet struct{}

func (t testEmptyCharSet) Characters() []rune {
	return []rune{}
}

func TestRandom(t *testing.T) {
	actual := Random()

	assert.Equals(t, DefaultRandomParameters.Length, uint(len(actual)))
	assertCharSet(t, actual, Alphanumeric{})
}

func TestRandomWithParameters_WithLength0(t *testing.T) {
	var parameters RandomParameters = RandomParameters{
		Length:       0,
		Prefix:       "",
		CharacterSet: Alphanumeric{},
	}

	actual := RandomWithParameters(parameters)

	assert.Equals(t, "", actual)
}

func TestRandomWithParameters_WithSpecifiedLength(t *testing.T) {
	var parameters RandomParameters = RandomParameters{
		Length:       30,
		Prefix:       "",
		CharacterSet: Alphanumeric{},
	}

	actual := RandomWithParameters(parameters)

	assert.Equals(t, 30, len(actual))
	assertCharSet(t, actual, Alphanumeric{})
}

func TestRandomWithParameters_WithNilCharacterSet(t *testing.T) {
	var parameters RandomParameters = RandomParameters{
		Length: 7,
		Prefix: "",
	}

	actual := RandomWithParameters(parameters)

	assert.Equals(t, parameters.Length, uint(len(actual)))
	assertCharSet(t, actual, Alphanumeric{})
}

func TestRandomWithParameters_WithSpecifiedCharSet(t *testing.T) {

	var parameters RandomParameters = RandomParameters{
		Length:       20,
		Prefix:       "",
		CharacterSet: testCharSet{},
	}

	actual := RandomWithParameters(parameters)

	assert.Equals(t, parameters.Length, uint(len(actual)))
	assertCharSet(t, actual, testCharSet{})
}

func TestRandomWithParameters_WithEmptyCharSet(t *testing.T) {

	var parameters RandomParameters = RandomParameters{
		Length:       20,
		Prefix:       "",
		CharacterSet: testEmptyCharSet{},
	}

	actual := RandomWithParameters(parameters)

	assert.Equals(t, 0, len(actual))
	assert.Equals(t, "", actual)
}

func assertCharSet(t *testing.T, actual string, charSet CharacterSet) {
	var chars []rune = charSet.Characters()
	for _, char := range actual {
		if !slices.Contains(chars, char) {
			t.Errorf("Unexpected character '%v' found in string", char)
		}
	}
}
