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
	fmt.Fscanln(reader, &N)
	str := ""
	for i := 0; i < N; i++ {
		vps := 0
		fmt.Fscanln(reader, &str)
		for _, c := range str {
			if c == '(' {
				vps++
			} else {
				vps--
				if vps < 0 {
					break
				}
			}
		}
		if vps == 0 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
