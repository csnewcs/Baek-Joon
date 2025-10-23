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
	nums := make([]int, 10)
	devide42 := make([]int, 42)
	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &nums[i])
		devide42[nums[i] % 42]++
	}
	different := 0
	for _, v := range(devide42) {
		if v != 0 {
			different++
		}
	}
	fmt.Fprint(writer, different)
}
