// 최소공배수: 60
package main

import "fmt"

func main() {
	var start, end int
	fmt.Scanln(&start)
	fmt.Scanln(&end)
	for ; start <= end; start += 60 {
		fmt.Printf("All positions change in year %d\n", start)
	}
}
