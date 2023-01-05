package main

import "fmt"

var num, min, max, count int

func main() {
	fmt.Print("Min number : ")
	fmt.Scanln(&min)
	fmt.Print("Max number : ")
	fmt.Scanln(&max)

	for num = min; num <= max; num++ {
		count := 0
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				count++
				break
			}
		}
		if count == 0 && num != 1 {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}
