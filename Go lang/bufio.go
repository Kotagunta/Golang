package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var res string
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	name, err := r.ReadString('\n')
	fmt.Print("Enter your address : ")
	add, _ := r.ReadString('\n')
	fmt.Print("Enter your ksddh : ")
	fmt.Scanln(&res)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s!\n", strings.TrimSpace(name))
	fmt.Printf("Hello %s!\n", strings.TrimSpace(add))
	fmt.Println(res)
}
