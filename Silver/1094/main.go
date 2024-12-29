package main

import "fmt"

func main() {
	var wantStickSize int
	fmt.Scan(&wantStickSize)
	count := 0
	for wantStickSize > 0 {
		if 1 & wantStickSize == 1 {
			count++
		}
		wantStickSize >>= 1
	}
	fmt.Println(count)
}