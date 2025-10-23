package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	var repeat int
	var str string
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &repeat, &str)
		p := ""
		for _, c := range(str) {
			p += strings.Repeat(string(c), repeat)
		}
		fmt.Fprintln(writer, p)
	}
}
