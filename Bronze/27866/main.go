package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int
var S string

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &S)
	fmt.Fscan(reader, &N)
	fmt.Fprint(writer, string(S[N - 1]))
}
