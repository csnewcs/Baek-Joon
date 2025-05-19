package main

import (
  "fmt"
  "math"
)

func main() {
  var K, N float64
  fmt.Scan(&N)
  fmt.Scan(&K)
  cherryPerWatermelon := math.Pow(2, N-1)
  max := int(K / cherryPerWatermelon)
  fmt.Print(max)
}