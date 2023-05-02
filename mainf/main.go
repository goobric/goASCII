package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	concatString := strings.Join(os.Args[1:], " ")
	fmt.Println(concatString)
	sepWords := strings.Split(concatString, `\n`)
	fmt.Println(sepWords)
}
