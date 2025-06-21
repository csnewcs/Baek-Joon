package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	p, q := (b - a), b //p: 분자, q: 분모
	for i := 1; i <= p; i++ {
		if p%i == 0 && q%i == 0 {
			p /= i
			q /= i
			i = 2
		}
	}
	fmt.Printf("%d %d", p, q)
}
