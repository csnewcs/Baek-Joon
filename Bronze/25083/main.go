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
	str := []string{
		`         ,r'"7`,
		"r`-_   ,'  ,/",
		` \. ". L_r'`,
		"   `~\\/",
		"      |",
		"      |",
	}
	for _, s := range str {
		fmt.Fprintln(writer, s)
	}
}
