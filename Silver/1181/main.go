package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	dict := make([]string, 0)
	var word string
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &word)
		if !slices.Contains(dict, word) {
			dict = append(dict, word)
		}
	}
	sort.Slice(dict, func(i, j int) bool {
		if len(dict[i]) != len(dict[j]) {
			return len(dict[i]) < len(dict[j])
		}
		for c := 0; true; c++ {
			if dict[i][c] != dict[j][c] {
				return dict[i][c] < dict[j][c]
			}
		}
		return false
	})
	for _, w := range dict {
		fmt.Fprintln(writer, w)
	}
}
