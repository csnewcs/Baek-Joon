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
	var input int
	sum := 0
	for i := 0; i < 5; i++ {
		fmt.Fscan(reader, &input)
		sum += input * input
	}
	fmt.Fprintln(writer, sum%10)
}
