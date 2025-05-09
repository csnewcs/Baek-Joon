package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var people []int
var people_2 []int

func main() {
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)
	people = make([]int, N)
	// people_2 = make([]int, N)
	tip := 0
	plus := 0
	// tip_2 := 0
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &people[i])
	}
	people = mergeSort(people)

	// used := make([]bool, N)
	for i := 0; i < N; i++ {
		if i >= people[i] {
			plus++
			continue
		}
		tip += people[i] - i + plus
	}
	fmt.Fprint(writer, tip)
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
