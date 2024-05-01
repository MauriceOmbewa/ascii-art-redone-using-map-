package main

import (
	printascii "ascii-map/ascii_map"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error with the number of arguments. Only two arguments required")
		return
	}

	arguments := os.Args[1]

	if arguments == "" {
		return
	} else if arguments == "\\n" {
		fmt.Println()
		return
	}

	for _, char := range arguments {
		if char < 32 || char > 126 {
			fmt.Println("Characters not within the required ascii characters")
			return
		}
	}

	inputfile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(inputfile)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	words := strings.Split(arguments, "\\n")

	printascii.AsciiArt(words, lines)
}
