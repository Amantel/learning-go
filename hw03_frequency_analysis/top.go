package hw03frequencyanalysis

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"
)

type occ struct {
	word    string
	occNumb int
}

func Top10(s string) []string {
	input := strings.TrimSpace(s)
	var frequentWords []string

	if len(input) == 0 {
		return frequentWords
	}
	words := make(map[string]int)
	var word string
	for _, char := range input {
		// fmt.Printf("Index: %d, Rune (character): %c, Type: %T\n", index, char, char)
		if unicode.IsSpace(char) {
			if word == "" || word == " " {
				continue
			}
			// fmt.Println("finished word", word)
			words[word]++
			word = ""
		} else {
			word += string(char)
		}
	}

	var occurances []occ
	for word, occNumb := range words {
		occurances = append(occurances, occ{word: word, occNumb: occNumb})
	}

	slices.SortFunc(occurances, func(a occ, b occ) int {
		if b.occNumb == a.occNumb {
			words := []string{a.word, b.word} // a - 0, b - 1
			sort.Strings(words)

			sortVal := 1
			if words[0] == a.word {
				sortVal = -1
			}

			return sortVal
		}
		return b.occNumb - a.occNumb
	})

	fmt.Println("post sorted occurances", occurances)

	for i, occurance := range occurances {
		if i > 9 {
			break
		}
		frequentWords = append(frequentWords, occurance.word)
	}

	return frequentWords
}
