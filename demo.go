package main

import (
	"fmt"
	//"strings"
)

func LongestCommonPrefix(strs [3]string) string {
	y := ""
	for i, j := range strs {
		_ = i
		y += j
	}
	return y

}

func main() {
	strs := [3]string{"flower", "flow", "flight"}
	fmt.Println(LongestCommonPrefix(strs))
}
