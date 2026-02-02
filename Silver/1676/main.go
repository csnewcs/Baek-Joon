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

	evenCount := N / 2
	fiveCount := 0

	for i := 1; i <= N; i++ {
		if i%5 == 0 {
			fiveCount++
			if i%25 == 0 {
				fiveCount++
				if i%125 == 0 {
					fiveCount++
				}
			}
		}
	}

	fmt.Fprintln(writer, min(evenCount, fiveCount))
}
