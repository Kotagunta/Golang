package main

import "fmt"

func anagram(str1 string, str2 string) string {
	if len(str1) != len(str2) {
		return "both the strings are not anagram"
	}
	dict := make(map[rune]int)
	for _, value := range str1 {
		dict[value]++
	}
	for _, value := range str2 {
		dict[value]--
	}
	for k, _ := range dict {
		if dict[k] != 0 {
			return "both the strings are not anagram"
		}
	}
	return "both the strings are anagram"
}

func main() {

	var word1, word2 string

	fmt.Print("enter a string1 : ")
	fmt.Scanln(&word1)
	fmt.Print("enter a string2 : ")
	fmt.Scanln(&word2)
	result := anagram(word1, word2)
	fmt.Println(result)
}
