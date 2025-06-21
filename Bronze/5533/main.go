package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var playerCount int
	fmt.Fscanln(reader, &playerCount)

	playerCards := make([][]int, playerCount)
	for i := 0; i < playerCount; i++ {
		playerCards[i] = make([]int, 3)
		fmt.Fscanln(reader, &playerCards[i][0], &playerCards[i][1], &playerCards[i][2])
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < playerCount-1; j++ {
			inspectCard := playerCards[j][i]
			for k := j + 1; k < playerCount; k++ {
				if inspectCard == playerCards[k][i] {
					playerCards[j][i] = 0
					playerCards[k][i] = 0
				}
			}
		}
	}
	for i := 0; i < playerCount; i++ {
		fmt.Fprintln(writer, playerCards[i][0]+playerCards[i][1]+playerCards[i][2])
	}
}
