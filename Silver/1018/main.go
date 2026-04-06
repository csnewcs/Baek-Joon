package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var M, N int
var whiteStartBoard = [8]uint8 {
	0b10101010,
	0b01010101,
	0b10101010,
	0b01010101,
	0b10101010,
	0b01010101,
	0b10101010,
	0b01010101,
}
var blackStartBoard = [8]uint8 {
	0b01010101,
	0b10101010,
	0b01010101,
	0b10101010,
	0b01010101,
	0b10101010,
	0b01010101,
	0b10101010,
}


func main() {
	defer writer.Flush()
	fmt.Fscanln(reader, &M, &N) //M: 세로, N: 가로
	//M*N
	board := make([]int64, M) //50*50이하
	var line string
	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &line)
		board[i] = parse(line)
	}
	least := 64
	tempBoard := [8]uint8{0}
	for i := 0; i < N-7; i++ {
		for j := 0; j < M-7; j++ {
			for l := 0; l < 8; l++ {
				for k := 0; k < 8; k++ {
					tempBoard[l] <<= 1
					tempBoard[l] += uint8((board[j+l] >> (i+k)) & 1)
				}
			}
			w, b := checkColorDifference(whiteStartBoard, tempBoard), checkColorDifference(blackStartBoard, tempBoard)
			if w < least {
				least = w
			}
			if b < least {
				least = b
			}
			if least == 0 {
				fmt.Fprint(writer, 0)
				return
			}
		}
		
	}
	fmt.Fprint(writer, least)
	//8*8 ~ 50*50 사이즈, 개별로 다 체크해본다면 1번~1849번 체크.
	//64*1849*2 = 236672
}
func parse(line string) int64 {
	var l int64
	for _, c := range line {
		l <<= 1
		if c == 'W' {
			l += 1
		}
	}
	return l
}
func checkColorDifference(board1 [8]uint8, board2 [8]uint8) int {
	count := 0
	for i, line := range board1 {
		xor := line^board2[i]
		for xor > 0 {
			if xor & 1 == 1 {
				count++
			}
			xor >>= 1
		}
	}
	return count
}