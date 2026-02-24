package hw03frequencyanalysis

import (
	"slices"
	"strings"
)

type occ struct {
	word    string
	occNumb int
}

func Top10(s string) []string {
	var frequentWords []string

	separatewords := strings.Fields(s)
	if len(separatewords) == 0 {
		return frequentWords
	}

	words := make(map[string]int)
	for _, word := range separatewords {
		words[word]++
	}

	var occurances []occ
	for word, occNumb := range words {
		occurances = append(occurances, occ{word: word, occNumb: occNumb})
	}

	slices.SortFunc(occurances, func(a occ, b occ) int {
		if b.occNumb == a.occNumb {
			if a.word < b.word {
				return -1
			}

			if a.word > b.word {
				return 1
			}

			return 0
		}
		return b.occNumb - a.occNumb
	})

	for i, occurance := range occurances {
		if i > 9 {
			break
		}
		frequentWords = append(frequentWords, occurance.word)
	}

	return frequentWords
}
