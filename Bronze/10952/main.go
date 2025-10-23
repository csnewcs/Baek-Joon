package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var A, B int
	fmt.Fscan(reader, &A, &B)
	for A != 0 && B != 0 {
		fmt.Fprintln(writer, A + B)
		fmt.Fscan(reader, &A, &B)
	}
}
