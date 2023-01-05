package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Go lang"
	string2 := "pyhton"
	string3 := "Go"

	//find the length of the string1
	str_length := len(string1)
	fmt.Println(str_length)

	// Add two strings
	fmt.Println(string1 + " " + string2)

	// Compare two strings
	fmt.Println(strings.Compare(string1, string2)) // return -1 because string1 is greater than string2.
	fmt.Println(strings.Compare(string2, string3)) // return 1 because string2 is smaller than string3.
	fmt.Println(strings.Compare(string1, string3)) // return 0 because string1 and string3 both are equal.

	// Check if string contains substring
	fmt.Println(strings.Contains(string1, string3)) //return true
	fmt.Println(strings.Contains(string2, string3)) //return false

	// Replace a string
	fmt.Println(strings.Replace(string1, "l", "L", 1))

	// Change case of string
	word1 := "msys"
	word2 := "TECHNOLOGIES"
	fmt.Println(strings.ToUpper(word1)) // return MSYS (convert lower case to upper case)
	fmt.Println(strings.ToLower(word2)) // return technologies (convert upper case to lower case)
}
