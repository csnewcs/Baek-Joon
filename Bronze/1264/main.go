package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for sentence, _ := input.ReadString('\n'); sentence[0] != '#'; sentence, _ = input.ReadString('\n') {
		vowelCount := 0
		for _, c := range sentence {
			if slices.Contains(vowels, c) {
				vowelCount++
			}
		}
		fmt.Println(vowelCount)
	}
}
