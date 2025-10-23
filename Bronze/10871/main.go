package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, X int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &X)
	numbers := make([]int, N)
	for i := range(numbers) {
		fmt.Fscan(reader, &numbers[i])
	}
	for _, num := range(numbers) {
		if num < X {
			fmt.Fprintf(writer, "%d ", num)
		}
	}
}
