package string_utils

import (
	"math/rand"
	"strings"
)

type RandomParameters struct {
	Length       uint
	Prefix       string
	CharacterSet CharacterSet
}

var DefaultRandomParameters RandomParameters = RandomParameters{
	Length:       15,
	CharacterSet: Alphanumeric{},
	Prefix:       "",
}

// Random creates a random string with the DefaultRandomParameters.
func Random() string {
	return RandomWithParameters(DefaultRandomParameters)
}

// RandomWithParameters creates a random string with the provided RandomParameters.
// If the CharacterSet is not specified, then the alphanumeric character set will be used.
func RandomWithParameters(parameters RandomParameters) string {
	var charSet CharacterSet = parameters.CharacterSet
	if charSet == nil {
		charSet = Alphanumeric{}
	}

	var chars []rune = charSet.Characters()
	if len(chars) == 0 {
		return ""
	}

	var builder strings.Builder
	for i := uint(0); i < parameters.Length; i++ {
		randomIndex := rand.Intn(len(chars))

		builder.WriteRune(chars[randomIndex])
	}

	return builder.String()
}
