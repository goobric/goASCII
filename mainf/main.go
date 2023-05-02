package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputString := strings.Join(os.Args[1:], " ")
	fmt.Println(inputString)
	sepWords := strings.Split(inputString, `\n`)
	fmt.Println(sepWords)
}
