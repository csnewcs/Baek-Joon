package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var A, B, C int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &A, &B, &C)
	count := make([]int, 10)
	for multiply := A * B * C; multiply > 0; multiply /= 10 {
		count[multiply % 10]++
	}
	for _, c := range(count) {
		fmt.Fprintln(writer, c)
	}
}
