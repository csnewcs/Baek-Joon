package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N int
var buildings []int
var visible [][]int

func main() {
	//A와 B를 잇는 일차함수가 있을 때, 그 사이에 다른 건물이 튀어나와있다면 안 보인다 판정
	defer writer.Flush()
	fmt.Fscan(reader, &N)
	if N == 1 {
		fmt.Fprint(writer, "0")
		return
	}
	buildings = make([]int, N)
	visible = make([][]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &buildings[i])
		visible[i] = make([]int, N)
	}
	max := -1
	for i := 0; i < N; i++ {
		count := 0
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			if visible[i][j] == 1{
				count++
				continue
			} else if visible[i][j] == -1 {
				continue
			}
			g := NewGraph(i, buildings[i], j, buildings[j])
			isVisible := true
			
			if j < i {
				for k := j + 1; k < i; k++ {
					if g.Calculate(k) <= float64(buildings[k]) {
						isVisible = false
						break
					}
				}
			} else {
				for k := i + 1; k < j; k++ {
					if g.Calculate(k) <= float64(buildings[k]) {
						isVisible = false
						break
					}
				}
			}

			if isVisible {
				count++
				visible[i][j] = 1
				visible[j][i] = 1
			} else {
				visible[i][j] = -1
				visible[j][i] = -1
			}
		}
		if max < count {
			max = count
		}
	}
	fmt.Fprint(writer, max)
}

type Graph struct {
	Slope     float64
	Intercept float64
}

func (g Graph) Calculate(x int) float64 {
	return g.Slope*float64(x) + g.Intercept
}
func NewGraph(x1 int, y1 int, x2 int, y2 int) Graph {
	slope := float64(y1-y2) / float64(x1-x2)
	return Graph{slope, float64(y1) - slope*float64(x1)}
}
