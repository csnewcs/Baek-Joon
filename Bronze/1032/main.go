package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var pattern string 
	fmt.Scan(&pattern)
	var name string
	for i := 1; i < N; i++ {
		fmt.Scan(&name)
		nameRune := []rune(name)
		for j, r := range(pattern) {
			if r != nameRune[j] {
				nameRune[j] = '?'
			}
		}
		pattern = string(nameRune)
	}
	fmt.Println(pattern)
}
