package main

import (
	"fmt"
	"unicode"
)

func clearRuneSlice(runeSlice []rune) []rune {
	var newRuneSlice []rune

	for _, runeValue := range runeSlice {
		if !unicode.IsSpace(runeValue) {
			newRuneSlice = append(newRuneSlice, runeValue)
		}
	}

	return newRuneSlice
}

func IsPalindromePhrase(s string) bool {
	stringLength := len(s)
	if stringLength == 0 || stringLength == 1 {
		return true
	}

	runeSlice := clearRuneSlice([]rune(s))
	runeSliceLength := len(runeSlice)
	hasEqulibrium := true

	for i, runeValue := range runeSlice {
		oppositeI := runeSliceLength - 1 - i
		oppositeRuneValue := runeSlice[oppositeI]
		if runeValue != oppositeRuneValue {
			hasEqulibrium = false
		}
	}

	return hasEqulibrium
}

func IsPalindrome(s string) bool {
	stringLength := len(s)
	if stringLength == 0 || stringLength == 1 {
		return true
	}
	runeSlice := []rune(s)
	runeSliceLength := len(runeSlice)
	hasEqulibrium := true

	for i, runeValue := range runeSlice {
		oppositeI := runeSliceLength - 1 - i
		oppositeRuneValue := runeSlice[oppositeI]
		if runeValue != oppositeRuneValue {
			hasEqulibrium = false
		}
	}

	return hasEqulibrium
}

func main() {
	fmt.Println(IsPalindrome(""))
	fmt.Println(IsPalindrome("111"))
	fmt.Println(IsPalindrome("abc"))
	fmt.Println(IsPalindrome("ccc"))
	fmt.Println(IsPalindrome("камак"))
	fmt.Println(IsPalindrome("llll"))
	fmt.Println(IsPalindrome("🧍🏽🧍🏽🧍🏽"))
}
