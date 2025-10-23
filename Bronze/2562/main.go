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
	numbers := make([]int, 9)
	for i := 0; i < 9; i++ {
		fmt.Fscan(reader, &numbers[i])
	}
	maxInt, maxIndex := numbers[0], 0
	for index, value := range numbers {
		if value > maxInt {
			maxInt = value
			maxIndex = index
		}
	}
	fmt.Fprintln(writer, maxInt)
	fmt.Fprintln(writer, maxIndex+1)
}
