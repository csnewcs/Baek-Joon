package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var N, M int //건물, 구름다리 수
	fmt.Fscan(reader, &N)
	fmt.Fscan(reader, &M)
	bridges := make([][]int, N)
	colors := make([]int, N)
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		bridges[x-1] = append(bridges[x-1], y-1)
		bridges[y-1] = append(bridges[y-1], x-1)
	}
	for i := 0; i < N; i++ {
		used := []int{}
		for _, connected := range bridges[i] {
			if colors[connected] != 0 {
				used = append(used, colors[connected])
			}
		}
		color := 1
		for slices.Contains(used, color) {
			color++
		}
		colors[i] = color
		fmt.Fprintf(writer, "%d ", color)
	}
}
