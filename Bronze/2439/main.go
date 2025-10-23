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
	for i := 1; i <= N; i++ {
		stars := fmt.Sprintf("%%%ds\n", N)
		fmt.Printf(stars, strings.Repeat("*", i))
	}
}
