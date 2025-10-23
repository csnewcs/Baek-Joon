package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var A, B, C int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &A, &B, &C)
	fmt.Fprintln(writer, A + B - C)

	str := fmt.Sprintf("%d%d", A, B)
	sum, _ := strconv.Atoi(str)

	fmt.Fprintln(writer, sum - C)
}
