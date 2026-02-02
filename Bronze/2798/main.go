package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, M int
var sumList []int = make([]int, 161700) //max C

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &M)
	numbers := make([]int, M)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &numbers[i])
	}
	index := 0
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				sumList[index] = numbers[i] + numbers[j] + numbers[k]
				index++
			}
		}
	}
	slices.Sort(sumList)
	for i := 161699; i > 1; i-- {
		if sumList[i] <= M {
			fmt.Fprintln(writer, sumList[i])
			return
		}
	}
}
