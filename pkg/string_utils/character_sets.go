package string_utils

import (
	"slices"
)

type CharacterSet interface {
	Characters() []rune
}

type Alphanumeric struct {
}

func (a Alphanumeric) Characters() []rune {
	var alphaNumeric []rune

	var lowerCaseLetters []rune = lowerCaseAlphabet()
	var upperCaseLetters []rune = upperCaseAlphabet()
	var numeric []rune = numeric()

	alphaNumeric = slices.Concat(alphaNumeric, lowerCaseLetters, upperCaseLetters, numeric)

	return alphaNumeric
}

func lowerCaseAlphabet() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
}

func upperCaseAlphabet() []rune {
	return []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
}

func numeric() []rune {
	return []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
}
