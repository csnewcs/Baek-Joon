package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var H, M int

func main() {
	defer writer.Flush()
	
	fmt.Fscan(reader, &H, &M)
	M -= 45
	if M < 0 {
		M += 60
		H -= 1
	}
	if H < 0 {
		H += 24
	}
	fmt.Fprint(writer, H, M)
}
