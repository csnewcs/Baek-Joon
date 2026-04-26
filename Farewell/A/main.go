package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscanln(reader, &N)
	// 2 -> 10 x 4
	// 3 -> 14 x 6
	// 4 -> 18 x 8?
	// 5 -> 22 x 10
	// A -> (4A + 2) x 2A

	// 라인당 별의 개수는 3개
	// 가장 아랫줄 기준 1번째, 끝에서 N번째, 끝에서 N+2번째 / 3N, 3N+2
	// 중간 기준 N번째, 끝에서 1번째, 끝에서 2N+1번째 / 4N+2, 2N -> 즉 이 둘을 합치면 언제나 6N+2 번째
	stars := make([][3]int, N * 2)
	for i := 0; i < N; i++ {
		stars[i][0] = i
		stars[i][2] = (3 * N + 2) + i
		stars[i][1] = (6 * N + 2) - stars[i][2]
	}
	for i := 0; i < N; i++ {
		stars[N+i][0] = N + i
		stars[N+i][1] = stars[N-i-1][1]
		stars[N+i][2] = stars[N-i-1][2]
	}
	for i := 2*N - 1; i > -1; i-- {
		str := strings.Repeat(" ", 4 * N + 2)
		str = str[:stars[i][0]] + "*" + str[stars[i][0]+1:stars[i][1]] + "*" + str[stars[i][1]+1:stars[i][2]] + "*" + str[stars[i][2]+1:]
		fmt.Fprintln(writer, str)
	}
}
