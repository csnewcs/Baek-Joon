package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int

const MOD int = 1000000000

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N) //길이, 1 <= N <= 100
	numberCount := make([][10]int, N + 1)
	for j := 1; j < 10; j++ {
		numberCount[1][j] = 1
	}
	for i := 1; i < N; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				numberCount[i+1][1] += numberCount[i][0] % MOD
				continue
			} else if j == 9 {
				numberCount[i+1][8] += numberCount[i][9] % MOD
				continue
			}
			numberCount[i+1][j-1] += numberCount[i][j] % MOD
			numberCount[i+1][j+1] += numberCount[i][j] % MOD
		}
	}
	
	count := 0
	for j := 0; j < 10; j++ {
		count += numberCount[N][j] % MOD
	}
	fmt.Fprint(writer, count % MOD)
}
