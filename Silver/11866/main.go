package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, K int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N, &K)
	people := make([]int, N)
	for i := range people {
		people[i] = i
	}
	lastRemoved := -1
	fmt.Fprint(writer, "<")
	for i := 0; i < N-1; i++ {
		lastRemoved = removePerson(K, &people, lastRemoved)
	}
	fmt.Fprint(writer, people[0]+1, ">")
}

func removePerson(k int, people *[]int, lastRemoved int) int {
	target := (lastRemoved + k) % len(*people)
	fmt.Fprint(writer, (*people)[target]+1, ", ")
	*people = slices.Concat((*people)[:target], (*people)[target+1:])
	return target-1
}