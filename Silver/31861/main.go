package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, M int // 0 <= N <= 25, 2 <= M <= 1000
const MOD int = 7 + 10e9
var canAfterAlphabet [26][]int = [26][]int{}
var alphabetStr string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	fmt.Fscan(reader, &M)
	//M자리 비밀번호, 알파벳 사이 거리는 N 이상
	for i := 0; i < 26; i++ {
		canAfterAlphabet[i] = make([]int, 0)
		for j := 0; j < 26; j++ {
			if i - j >= N {
				canAfterAlphabet[i] = append(canAfterAlphabet[i], j)
			} else if j - i >= N {
				canAfterAlphabet[i] = append(canAfterAlphabet[i], j)
			}
		}
	}

	sum := 0
	for i := 0; i < 26; i++ {
		if len(canAfterAlphabet[i]) == 0 {
			continue
		}
		sum += (getPasswordCount(i, M-1) % MOD)
	}
	fmt.Fprint(writer, sum)
}
func getPasswordCount(startAlphabet, length int) int {
	if length == 1 {
		return len(canAfterAlphabet[startAlphabet])
	}
	sum := 0
	for _, i := range(canAfterAlphabet[startAlphabet]) {
		sum += getPasswordCount(i, length - 1)
		sum %= MOD
	}
	return sum
}