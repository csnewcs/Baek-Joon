package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var T int
var N int

const DEVIDE int = 1000000007

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(reader, &N)
		fmt.Fprintln(writer, f(N+1))
	}
}

var memory []int = make([]int, 191231)

func f(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memory[n] != 0 {
		return memory[n]
	}
	memory[n] = (f(n-1) + f(n-2)) % DEVIDE
	return memory[n]
}
