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

	var N int
	fmt.Fscanln(reader, &N)

	k := make([]int, 2000001)
	var num int
	fmt.Fscanln(reader, &num)
	k[num+1000000]++
	for i := 1; i < N; i++ {
		fmt.Fscanln(reader, &num)
		k[num+1000000]++
		// if max < num {
		// 	max = num
		// } else if min > num {
		// 	min = num
		// }
	}
	for i := 0; i < 2000001; i++ {
		print := i - 1000000
		for j := 0; j < k[i]; j++ {
			fmt.Fprintln(writer, print)
		}
	}
}
