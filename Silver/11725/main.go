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
	fmt.Fscan(reader, &N)

	adj := make([][]int, N+1)
	for i := 1; i < N; i++ {
		var x, y int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}

	parent := make([]int, N+1)
	visited := make([]bool, N+1)

	queue := []int{1}
	visited[1] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range adj[current] {
			if !visited[child] {
				parent[child] = current
				visited[child] = true
				queue = append(queue, child)
			}
		}
	}

	for i := 2; i <= N; i++ {
		fmt.Fprintln(writer, parent[i])
	}
}
