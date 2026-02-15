package main

import (
	"fmt"
	"unicode"
)

func InspectPhrase(s string) (char rune, count int, isComplex bool) {
	fmt.Println(s)

	runeSlice := []rune(s)

	myMap := make(map[rune]int)
	var charFrequency int
	var mostFrequentChar rune
	var validCharsLen int

	for _, runeVal := range runeSlice {
		if !unicode.IsLetter(runeVal) {
			continue
		}
		validCharsLen++

		myMap[runeVal]++
		localFrequency := myMap[runeVal]
		if localFrequency > charFrequency {
			charFrequency = localFrequency
			mostFrequentChar = runeVal
		}

	}

	fmt.Println(charFrequency, mostFrequentChar, myMap)

	if validCharsLen == 0 {
		return ' ', 0, false
	}

	if charFrequency >= validCharsLen/4 {
		return mostFrequentChar, charFrequency, true
	}

	return ' ', 0, false

}

func main() {
	fmt.Println(InspectPhrase("«Во дворе трава на траве дрова»"))
	fmt.Println(InspectPhrase("Карл у Клары украл кораллы. А Клара у Карла украла кларнет"))
	fmt.Println(InspectPhrase(""))
	fmt.Println(InspectPhrase("пппппп"))
	fmt.Println(InspectPhrase("Интрига панов Пшекшицюльского и Кшепшицюльского"))
	fmt.Println(InspectPhrase("Жили ли ежи? Ежели ежи жили, ели ли ежи жужелиц?"))
	fmt.Println(InspectPhrase("топот"))
	fmt.Println(InspectPhrase("абв, абв, о"))
	fmt.Println(InspectPhrase("абв, абв, о"))
	fmt.Println(InspectPhrase("гггггоооо"))
	fmt.Println(InspectPhrase("гггггoooooo"))
	fmt.Println(InspectPhrase("     "))

}
