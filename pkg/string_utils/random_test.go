package string_utils

import (
	"slices"
	"strings"
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

func TestRandomWithPrefix(t *testing.T) {
	var prefix string = "prefix-"

	actual := RandomWithPrefix(prefix)

	assert.Equals(t, DefaultRandomParameters.Length, uint(len(actual)))
	assert.Equals(t, true, strings.HasPrefix(actual, prefix))
}

func TestRandomWithPrefix_WithLongPrefix(t *testing.T) {
	var size int = 100
	var builder strings.Builder
	for i := 0; i < size; i++ {
		builder.WriteRune('a')
	}

	prefix := builder.String()

	actual := RandomWithPrefix(prefix)

	assert.Equals(t, size, len(actual))
	assert.Equals(t, prefix, actual)
}

func assertCharSet(t *testing.T, actual string, charSet CharacterSet) {
	var chars []rune = charSet.Characters()
	for _, char := range actual {
		if !slices.Contains(chars, char) {
			t.Errorf("Unexpected character '%v' found in string", char)
		}
	}
}
