package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	var A, B int
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &A)
		fmt.Fscan(reader, &B)
		fmt.Fprintln(writer, fastMod(A, B, 1, 10))
	}
}

func fastMod(base int, ex int, lastResult int, mod int) int {
	if ex&1 == 1 {
		lastResult = (lastResult * base) % mod
	}
	base = (base * base) % 100
	ex /= 2
	if ex == 0 {
		if lastResult == 0 {
			return 10
		}
		return lastResult
	}
	return fastMod(base, ex, lastResult, mod)
}
