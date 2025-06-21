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

	var N, K uint32 //N: 수빈, K: 동생(부동), 예시 입력: 5 17, 예시 출력: 4
	fmt.Fscanln(reader, &N, &K)

	nodes := make([]uint32, 0)
	nextNodes := [3]uint32{}
	depth := make(map[uint32]uint16)

	nodes = append(nodes, N)
	depth[N] = 0
	for len(nodes) != 0 {
		now := nodes[0]
		nodes = nodes[1:]
		if now == K {
			depth[K] = depth[now]
			break
		}

		for _, next := range nextNodes {
			if depth[next] == 0 {
				nodes = append(nodes, next)
				depth[next] = depth[now] + 1
			}
		}
	}
	fmt.Fprintln(writer, depth[K])
}
