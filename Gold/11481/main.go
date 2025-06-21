package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var N int
	var G float64
	fmt.Fscanf(reader, "%d %f", &N, &G)

	D := make([]int, N)
	theta := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &D[i])
        fmt.Fscan(reader, &theta[i])
	}

	var startSpeed float64 = 0
	var rad float64
    results := make([]float64, N)
	for i := N-1; i > -1; i-- {
		rad = float64(theta[i]) * math.Pi / 180
		a := G * math.Cos(rad)
		t := ((-1 * startSpeed) + math.Sqrt(startSpeed*startSpeed+(2*a*float64(D[i])))) / a
		v := a * t + startSpeed
		startSpeed = v
        results[i] = v
	}
    for i := 0; i < N; i++ {
		fmt.Fprintln(writer, results[i])
    }
}
