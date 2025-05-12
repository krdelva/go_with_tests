package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	/*
		for i := 0; i < 5; i++ {
			repeated += character
		}
	*/
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
