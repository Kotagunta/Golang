package main

import "fmt"

var total int = 0

func SumOfDigits(num int) int {
	if num > 0 {
		total += num % 10
		SumOfDigits(num / 10)
	}
	return total
}

func main() {
	var num int = 0
	var result int = 0
	fmt.Printf("Number: ")
	fmt.Scanf("%d", &num)

	result = SumOfDigits(num)

	fmt.Printf("%d\n", result)
}
