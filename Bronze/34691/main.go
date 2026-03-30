package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N string

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	for N != "end" {
		switch N {
		case "animal":
			fmt.Fprintln(writer, "Panthera tigris")
		case "flower":
			fmt.Fprintln(writer, "Forsythia koreana")
		case "tree":
			fmt.Fprintln(writer, "Pinus densiflora")
		}
		fmt.Fscan(reader, &N)
	}
}
