package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	if N == 0 {
		fmt.Fprint(writer, 0)
		return
	}
	opinions := getOpinions(N)
	afterCutEdge := cutEdge(N, opinions)
	sum := 0
	for _, opinion := range afterCutEdge {
		sum += opinion
	}
	fmt.Fprint(writer, (sum+len(afterCutEdge)/2)/len(afterCutEdge))
}

func getOpinions(n int) []int {
	opinions := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &opinions[i])
	}
	return opinions
}

func cutEdge(n int, opinions []int) []int {
	cutRange := (n*15 + 50) / 100
	slices.Sort(opinions)
	afterCut := opinions[cutRange : n-cutRange]
	return afterCut
}
