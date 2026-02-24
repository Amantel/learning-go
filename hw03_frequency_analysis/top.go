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
	// fmt.Println("pre sorted occurances", occurances)
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

	minNumber := occurances[9].occNumb
	fmt.Println("post sorted occurances", occurances)

	for i, occurance := range occurances {
		if occurance.occNumb > minNumber && i < 9 {
			frequentWords = append(frequentWords, occurance.word)
			continue
		}
		// надо ещё хвост лексикографически отсортировать
		bestLastWord := getBestLastWord(i, minNumber, occurances)
		frequentWords = append(frequentWords, bestLastWord...)
		return frequentWords

	}

	return frequentWords
}

func getBestLastWord(lastIndex int, minNumber int, occurances []occ) []string {
	var tail []string
	for i := lastIndex; true; i++ {
		if occurances[i].occNumb != minNumber {
			break
		}
		tail = append(tail, occurances[i].word)
	}

	// отсортировать лексикографически хвост
	fmt.Println("tail", tail)
	sort.Strings(tail)
	fmt.Println("tail", tail)
	return tail[0 : 10-lastIndex]
}
