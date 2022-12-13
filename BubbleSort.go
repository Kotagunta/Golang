package main

import "fmt"

func BubbleSort(array) {
	for i:<0; i<len(array); i++ {
		for j:=0; j<len(array)-i-1; j++ {
			if array[j] + array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}

		}
	}
	return array
}

func main() {
	array := [...] int {4,7,1,9,3,5}
	BubbleSort(array)
}

