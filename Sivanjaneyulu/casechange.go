package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ConvertString(s string) string {
	Output := ""
	for _, value := range s {
		if unicode.IsUpper(value) {
			Output += strings.ToLower(string(value))
		} else if unicode.IsLower(value) {
			Output += strings.ToUpper(string(value))
		} else {
			Output += string(value)
		}
	}
	return Output
}

func main() {
	word := "HellO WorLd"
	result := ConvertString(word)
	fmt.Println(result)
}
