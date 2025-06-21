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

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	fmt.Fscan(reader, &M)
	numbers := make(map[int]bool)
	var number int
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &number)
		numbers[number] = true
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(reader, &number)
		numbers[number] = false
	}
	res := make([]int, 0)
	for num, notDuplicated := range(numbers) {
		if notDuplicated {
			res = append(res, num)
		}
	}
	fmt.Fprintln(writer, len(res))
	slices.Sort(res)
	for _, num := range(res) {
		fmt.Fprintf(writer, "%d ", num)
	}
}
