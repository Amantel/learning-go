package main

import (
	"fmt"
	"unicode"
)

func fixSlice(runeSlice []rune) []rune {
	var newRuneSlice []rune

	for _, runeValue := range runeSlice {
		if unicode.IsLetter(runeValue) {
			newRuneSlice = append(newRuneSlice, unicode.ToLower(runeValue))
		}
	}

	return newRuneSlice
}

func InspectPhrase(s string) (char rune, count int, isComplex bool) {
	fmt.Println(s)
	runeSlice := []rune(s)
	fixedSlice := fixSlice(runeSlice)
	maxTargetCount := len(fixedSlice)/4 + 1

	myMap := make(map[rune]int)

	for _, runeVal := range fixedSlice {
		myMap[runeVal]++

		if num+1 >= maxTargetCount {
			isComplex = true
			count = num + 1
			char = runeVal
		}
	}

	fmt.Println(maxTargetCount, myMap)

	return char, count, isComplex
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

}
