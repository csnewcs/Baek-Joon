package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)

	numberCounts := make(map[int]int) //-4000 ~ 4000
	scan := 0
	sum := 0
	minimum := math.MaxInt64
	maximum := math.MinInt64
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &scan)
		sum = sum + scan
		numberCounts[scan+4000]++
		minimum = min(minimum, scan)
		maximum = max(maximum, scan)
	}
	avg := int(math.Round(float64(sum) / float64(N)))
	fmt.Fprintln(writer, avg)
	midIndex := (N + 1) / 2
	checked := 0
	modes := make([]int, 1)
	modeCount := 0
	for num := minimum; num <= maximum; num++ {
		checked += numberCounts[num+4000]
		if checked >= midIndex {
			fmt.Fprintln(writer, num)
			checked -= 500000
		}
		if modeCount == numberCounts[num+4000] {
			modes = append(modes, num)
		} else if modeCount < numberCounts[num+4000] {
			modes = modes[:1]
			modeCount = numberCounts[num+4000]
			modes[0] = num
		}
	}
	if len(modes) != 1 {
		fmt.Fprintln(writer, modes[1])
	} else {
		fmt.Fprintln(writer, modes[0])
	}
	fmt.Fprintln(writer, maximum-minimum)
}
