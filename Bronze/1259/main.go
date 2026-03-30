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
	for N != 0 {
		if N < 10 {
			fmt.Fprintln(writer, "yes")
		} else if N < 100 {
			if N/10 == N%10 {
				fmt.Fprintln(writer, "yes")
			} else {
				fmt.Fprintln(writer, "no")
			}
		} else if N < 1000 {
			if N/100 == N%10 {
				fmt.Fprintln(writer, "yes")
			} else {
				fmt.Fprintln(writer, "no")
			}
		} else if N < 10000 {
			if (N/1000) == (N%10) && ((N/100)%10) == ((N%100)/10) {
				fmt.Fprintln(writer, "yes")
			} else {
				fmt.Fprintln(writer, "no")
			}
		} else {
			if (N/10000) == (N%10) && ((N/1000)%10) == ((N%100)/10) {
				fmt.Fprintln(writer, "yes")
			} else {
				fmt.Fprintln(writer, "no")
			}
		}
		fmt.Fscan(reader, &N)
	}
}
