package main

import (
	"fmt"
	"math"
)

var loop int = 0
var distances []int = make([]int, 0)
var notDuplicatedPoints string = ""

func main() {
	var T int //테스트 케이스의 개수
	fmt.Scan(&T)
	var N int //점의 개수, 짝수
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		var x, y int
		points := make([]Point, N)
		for j := 0; j < N; j++ {
			fmt.Scanf("%d %d", &x, &y)
			points[j] = Point{x, y}
		}
		distance := getShortest(points, N)
		fmt.Println(math.Sqrt(float64(distance)))
	}
}

type Point struct {
	X int
	Y int
}
type Vector struct {
	Start *Point
	End   *Point
	X     int
	Y     int
}

func factorial(n int) int {
	result := 1
	for i := 2; i < n; i++ {
		result = result * i
	}
	return result * n
}

func getShortest(points []Point, length int) int {
	cases := 1 << length
	shortest := math.MaxInt64
	for i := 0; i < cases; i++ { //벡터 수 = N/2 -> 경우의 수: 2^N
		if countTrue(i) == length/2 { //절반: plusPoint, 절반: minusPoint
			plusPoint := Point{0, 0}  //vector = P2 - P1
			minusPoint := Point{0, 0} // -> vector + vector = (P2 + P4) - (P1 + P3) -> plusPoint - minusPoint = totalPoint
			for j := 0; j < length; j++ {
				if (i & (1 << j)) != 0 {
					plusPoint.Add(points[j])
				} else {
					minusPoint.Add(points[j])
				}
			}
			totalPoint := Point{plusPoint.X - minusPoint.X, plusPoint.Y - minusPoint.Y}
			distance := totalPoint.GetDistance()
			if shortest > distance {
				shortest = distance
			}
		}
	}
	return shortest
}

func (a *Point) Add(p Point) {
	a.X = a.X + p.X
	a.Y = a.Y + p.Y
}
func (v *Point) GetDistance() int {
	return v.X*v.X + v.Y*v.Y
}
func countTrue(packedBool int) int {
	count := 0
	for packedBool > 0 {
		count = count + (packedBool & 1)
		packedBool = packedBool >> 1
	}
	return count
}