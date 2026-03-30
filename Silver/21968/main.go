package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var T int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		sum := 0
		var N int
		fmt.Fscan(reader, &N)
		for mask, sup := 1, 1; mask <= N; mask <<= 1 {
			if mask & N != 0 {
				sum += sup
			}
			sup *= 3
		}
		fmt.Fprintln(writer, sum)
	}
}
