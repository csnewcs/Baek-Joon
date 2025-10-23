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
	fmt.Fscan(reader, &A)
	fmt.Fscan(reader, &B)
	if A > B {
		fmt.Fprintln(writer, ">")
	} else if A < B {
		fmt.Fprintln(writer, "<")
	} else {
		fmt.Fprintln(writer, "==")
	}
}
