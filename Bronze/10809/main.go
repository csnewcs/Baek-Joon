package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var word string

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &word)
	indexes := make([]int, 26)
	for i, c := range(word) {
		if indexes[c - 'a'] == 0 {
			indexes[c - 'a'] = i + 1
		}
	}
	for _, i := range(indexes) {
		fmt.Fprint(writer, i-1, " ")
	}
}
