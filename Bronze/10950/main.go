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
		fmt.Fscan(reader, &A, &B)
		fmt.Fprintln(writer, A + B)
	}
}
