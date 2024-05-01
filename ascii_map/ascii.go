package printascii

import "fmt"

func AsciiArt(words []string, lines []string) {
	for _, word := range words {
		if word != "" {
			ascii := make(map[rune]string)
			for i := 0; i < 8; i++ {
				for _, char := range word {
					ascii[char] = lines[int(char-' ')*9+1+i]
					fmt.Print(ascii[char])

				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}
