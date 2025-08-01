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

// Random creates a random string with the string_utils.DefaultRandomParameters.
func Random() string {
	return RandomWithParameters(DefaultRandomParameters)
}

// RandomWithPrefix creates a random string that starts with the provided prefix.
// The rest of the random string attributes will be dictated by string_utils.DefaultRandomParameters.
func RandomWithPrefix(prefix string) string {
	params := DefaultRandomParameters
	params.Prefix = prefix

	return RandomWithParameters(params)
}

// RandomWithParameters creates a random string with the provided RandomParameters.
// If the CharacterSet is not specified, then the alphanumeric character set will be used.
// If the prefix is longer than the length then the prefix will be returned.
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
	prefixSize := len(parameters.Prefix)
	if prefixSize != 0 {
		if prefixSize >= int(parameters.Length) {
			return parameters.Prefix
		}

		builder.WriteString(parameters.Prefix)
	}

	for i := 0; i < int(parameters.Length)-prefixSize; i++ {
		randomIndex := rand.Intn(len(chars))
		builder.WriteRune(chars[randomIndex])
	}

	return builder.String()
}
