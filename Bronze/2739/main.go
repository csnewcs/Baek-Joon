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
	for i := 1; i < 10; i++ {
		fmt.Fprintf(writer, "%d * %d = %d\n", N, i, N*i)
	}
}
