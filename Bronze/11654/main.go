package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N string

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	r := byte(N[0])
	fmt.Fprintln(writer, r)
}
