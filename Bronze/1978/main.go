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
	primeNums := makePrimeNumberList(1000)
	fmt.Fscan(reader, &N)
	number := 0
	count := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &number)
		if slices.Contains(primeNums, number) {
			count++
		}
	}
	fmt.Fprint(writer, count)
}

func makePrimeNumberList(max int) []int {
	primeNumbers := make([]int, 4)
	primeNumbers[0] = 2
	primeNumbers[1] = 3
	primeNumbers[2] = 5
	primeNumbers[3] = 7 //10이하에서 소수들
	for i := 10; i < max; i++ {
		isPrimeNumber := true
		for _, primeNumber := range(primeNumbers) {
			if i % primeNumber == 0 {
				isPrimeNumber = false
				break
			}
		}
		if isPrimeNumber {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}