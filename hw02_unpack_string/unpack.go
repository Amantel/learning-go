package hw02unpackstring

import (
	"errors"
	"fmt"

	// "strconv"
	// "strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// func Unpack(s string) (string, error) {
// 	fmt.Println(s, "========>")
// 	if s == "" {
// 		return "", nil
// 	}

// 	// var myStr []rune
// 	// for _, r := range s {
// 	// 	if r != rune('\\') {
// 	// 		myStr = append(myStr, r)
// 	// 	}
// 	// }
// 	myStr := []rune(s)
// 	fmt.Println(string(myStr), "========>")
// 	var curr, next rune
// 	var finalString []rune
// 	for i := 0; i < len(myStr); i++ {
// 		if i == len(myStr)-1 {
// 			finalString = append(finalString, myStr[i])
// 			continue
// 		}
// 		curr = myStr[i]
// 		next = myStr[i+1]
// 		fmt.Println("---", string(curr), string(next), !unicode.IsNumber(curr), unicode.IsNumber(next))
// 		if !unicode.IsNumber(curr) && unicode.IsNumber(next) {
// 			i++ // jump to next rune after number
// 			for j := 0; j < int(next-'0'); j++ {
// 				finalString = append(finalString, curr)
// 			}
// 			fmt.Println(i, int(next-'0'), string(finalString))
// 		} else if !unicode.IsNumber(curr) && !unicode.IsNumber(next) {
// 			finalString = append(finalString, curr)
// 		} else {
// 			fmt.Println("RETURN ErrInvalidString")
// 			return "", ErrInvalidString
// 		}
// 	}

// 	fmt.Println(string(finalString))
// 	return string(finalString), nil
// }

func Unpack(s string) (string, error) {
	fmt.Println(s, "========>")
	if s == "" {
		return "", nil
	}

	myStr := []rune(s)
	fmt.Println(string(myStr), "========>")
	var curr, next rune
	var finalString []rune
	for i := 0; i < len(myStr); i++ {
		if i == len(myStr)-1 {
			finalString = append(finalString, myStr[i])
			continue
		}
		curr = myStr[i]
		next = myStr[i+1]
		fmt.Println("---", string(curr), string(next), !unicode.IsNumber(curr), unicode.IsNumber(next))

		if curr == '\\' {
			curr = next
			finalString = append(finalString, curr)
			fmt.Println("Appeded quoted number")
			i++

			fmt.Println("Nums", i+1, len(myStr), i+1 >= len(myStr))
			if i+1 >= len(myStr) {
				// избегаем выхода за рамки массива
				continue
			}

			next = myStr[i+1]
			fmt.Println("Next--->", string(next))
			if unicode.IsNumber(next) {
				fmt.Println("More additions", int(myStr[i+1]-'0'-1))
				for j := 0; j < int(myStr[i+1]-'0'-1); j++ {
					finalString = append(finalString, curr)
				}
				i++

			}

			continue
		}

		if !unicode.IsNumber(curr) && unicode.IsNumber(next) {
			i++ // jump to next rune after number
			for j := 0; j < int(next-'0'); j++ {
				finalString = append(finalString, curr)
			}
			fmt.Println(i, int(next-'0'), string(finalString))
		} else if !unicode.IsNumber(curr) && !unicode.IsNumber(next) {
			finalString = append(finalString, curr)
		} else {
			fmt.Println("RETURN ErrInvalidString")
			return "", ErrInvalidString
		}
	}

	fmt.Println(string(finalString))
	return string(finalString), nil
}
