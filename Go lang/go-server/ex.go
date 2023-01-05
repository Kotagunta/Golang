package main

import (
	"fmt"
	"time"
)

func msg(msg string) string {
	return msg
}

func main() {
	go fmt.Println(msg("this is thread"))
	time.Sleep(5)
}
