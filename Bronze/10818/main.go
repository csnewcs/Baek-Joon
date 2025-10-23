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
	var num int
	fmt.Fscan(reader, &num)
	min := num
	max := num
	for i := 1; i < N; i++ {
		fmt.Fscan(reader, &num)
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	fmt.Fprint(writer, min, max)
}
