package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var a, b, c int

func main() {
	defer writer.Flush()
	for true {
		fmt.Fscan(reader, &a, &b, &c)
		if a == b && b == c && c == 0 {
			return
		}
	}
}
