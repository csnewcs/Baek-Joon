package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, K int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &N) //N일 후
	fmt.Fscan(reader, &K) //3번 꺼질때마다 지연되는 시간
	const MINUTETODAY int = 24*60
	waitTerm := []int{3*60, 3*60, 18*60 + K} //15시 시작이면 15 -> 18 -> 21 -> 15. 3, 3, 18시간 간격
	currentTime := 15 * 60
	stoppedCount := 0
	countDay := N * MINUTETODAY
	countAfterDay := (N+1) *MINUTETODAY

	printTimes := make([]string, 0)
	for currentTime < countAfterDay {
		if currentTime >= countDay {
			printTimes = append(printTimes, fmt.Sprintf("%02d:%02d", (currentTime % MINUTETODAY) / 60, (currentTime % MINUTETODAY) % 60))
		}
		currentTime += waitTerm[stoppedCount % 3]
		stoppedCount++
	}
	fmt.Fprintln(writer, len(printTimes))
	for _, printTime := range(printTimes) {
		fmt.Fprintln(writer, printTime)
	}
}