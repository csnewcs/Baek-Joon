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
	_, err := fmt.Fscan(reader, &A, &B)
	for ; err == nil; _, err = fmt.Fscan(reader, &A, &B) {
		fmt.Fprintln(writer, A + B)
	}
}
