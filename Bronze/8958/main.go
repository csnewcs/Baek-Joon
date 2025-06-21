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

	for i := 0; i < N; i++ {
		var ox string
		fmt.Fscan(reader, &ox)
		score := 0
		plus := 1
		for j := 0; j < len(ox); j++ {
			if ox[j] == 'O' {
				score += plus
				plus++
			} else {
				plus = 1
			}
		}
		fmt.Fprintln(writer, score)
	}
}
