package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var A, B int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &A, &B)
	fmt.Fprintln(writer, A+B)
	fmt.Fprintln(writer, A-B)
	fmt.Fprintln(writer, A*B)
	fmt.Fprintln(writer, A/B)
	fmt.Fprintln(writer, A%B)
}
