package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	// 	(a+b+c+d+M)/5 (원래 평균)
	// -> ((a+b+c+d+M)/M)/5*100 (새로운 평균)
	score := 0
	sumScore := 0
	biggestScore := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &score)
		if score > biggestScore {
			biggestScore = score
		}
		sumScore += score
	}
	newAvg := float64(sumScore) / float64(N * biggestScore) * 100
	fmt.Fprint(writer, newAvg)
}
