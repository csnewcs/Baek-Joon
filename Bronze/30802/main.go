package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, T, P int

func main() {
	defer writer.Flush()
	shirtSizes := make([]int, 6)
	fmt.Fscan(reader, &N)
	for i := 0; i < 6; i++ {
		fmt.Fscan(reader, &shirtSizes[i])
	}
	fmt.Fscan(reader, &T, &P)
	shirtSum := 0
	for i := 0; i < 6; i++ {
		shirtSum += shirtSizes[i] / T + min((math.MaxInt & (shirtSizes[i] % T)), 1)
	}
	fmt.Fprintln(writer, shirtSum)
	fmt.Fprint(writer, N / P, N % P)
}