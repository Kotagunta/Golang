package main

import "fmt"

func main() {
	var number, temp, remainder int
	var result int = 0
	fmt.Println("Enter the number ..! : ")
	fmt.Scanf("%d", &number)
	//number = 153
	temp = number
	// Use of For Loop
	for number > 0 {
		remainder = number % 10
		result += remainder * remainder * remainder
		number /= 10
	}
	if result == temp {
		fmt.Printf("%d is an Armstrong number.", temp)
	} else {
		fmt.Printf("%d is not an Armstrong number.", temp)
	}
}
