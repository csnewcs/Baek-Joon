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
	switch N / 10 {
	case 9, 10:
		fmt.Fprintln(writer, "A")
	case 8:
		fmt.Fprintln(writer, "B")
	case 7:
		fmt.Fprintln(writer, "C")
	case 6:
		fmt.Fprintln(writer, "D")
	default:
		fmt.Fprintln(writer, "F")
	}
}
