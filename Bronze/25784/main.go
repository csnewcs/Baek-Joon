package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a, b, c int
	fmt.Fscanf(reader, "%d %d %d", &a, &b, &c)

	if a > b {
		if c > a { //c is biggest
			fmt.Fprint(writer, findRelationship(a, b, c))
		} else { //a is biggest
			fmt.Fprint(writer, findRelationship(b, c, a))
		}
	} else {
		if c > b { //c is biggest
			fmt.Fprint(writer, findRelationship(a, b, c))
		} else { //b is biggest
			fmt.Fprint(writer, findRelationship(a, c, b))
		}
	}
}

func findRelationship(a int, b int, c int) int {
	if isSum(a, b, c) {
		return 1
	} else if isProduct(a, b, c) {
		return 2
	}
	return 3
}
func isSum(a int, b int, c int) bool {
	return c == a+b
}
func isProduct(a int, b int, c int) bool {
	return c == a*b
}
