package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// concatenate the arguments from index[1] to one string
	concatString := strings.Join(os.Args[1:], " ")
	fmt.Println(concatString)
	// seperate the characters to a new line when \n is encountered
	sepWords := strings.Split(concatString, `\n`)
	fmt.Println(sepWords)
	// error handling
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	// count number of lines in each string length
	lines := strings.Split(string(file), "\n")
	// count number of words in each line
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				fmt.Println()
		}
		continue
	}
}
