package main

import "fmt"

func duplicate(arr []int) int {
	visitnumber := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visitnumber[arr[i]] == true {
			return arr[i]
		} else {
			visitnumber[arr[i]] = true
		}
	}
	return 0
}
func missing()

func main() {
	array1 := []int{17, 19, 18, 22, 21}
	duplicateNumber := duplicate(array1)
	missingNumber := missing(array1)
	fmt.Println(result)
}
