package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string as an argument")
		return
	}

	input := os.Args[1]

	if len(input) == 0 {
		fmt.Println("The input string is empty")
		return
	}

	for _, char := range input {
		fmt.Printf("%c ", char)
	}
	fmt.Println()
}

/* func main() {
	if len(strings.TrimSpace(os.Args[1])) == 0 {
		fmt.Println("Please provide a string as an argument")
		return
	}

	// ASCII representation of the input string
	ascii := [][]string{
		{" _  _  _  _  _  _  _  _  _ ", "| || || || || || || || || |", "|_||_||_||_||_||_||_||_||_|"},
		{"    _  _    ", "  | _| _|   |", "  ||_  _|   |"},
		{" _  _  _  _ ", " _| _| _| _|", "|_ |_  |_  |"},
		{" _        _ ", "|_|| || || |", " _| ||_||_||"},
		{" _  _     _ ", " _| _|  ||_ ", "|_  _|  | _)"},
	}

	// Convert the input string to a list of digits
	digits := make([]int, 0)
	for _, char := range os.Args[1] {
		digits = append(digits, int(char-'0'))
	}

	// Print the ASCII representation of the input string
	for row := range ascii {
		for _, digit := range digits {
			fmt.Print(ascii[row][digit])
		}
		fmt.Println()
	}
} */

// printArt is a helper function that takes a rune array and an array of lines containing the art data
// and prints the ascii art representation of the input string

/* func printArt(arr []rune, lines []string) {
	// iterate throughthe 8 lines of ascii art for each character
	if len(arr) != 0 {
		for line := 1; line <= 8; line++ {
			// loop through each character (rune) in the array slice
			for _, r := range arr {
				// calc the index offset for the current character's ascii art
				// each character has 9 lines of ascii art (including an empty line for separation)
				// use the Skip Value to find the correct starting index for the current character art
				skip := (r - 32) * 9

				// print the corresponding line of ascii art for the current character
				time.Sleep(300 * time.Millisecond)
				fmt.Print(lines[line+int(skip)])
			}
			time.Sleep(300 * time.Millisecond)
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
} */
