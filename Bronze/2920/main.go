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
	nums := make([]int, 8)
	mod := 0
	modString := [3]string{
		"ascending",
		"descending",
		"mixed",
	}

	for i := 0; i < 8; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	if nums[0] > nums[1] {
		mod = 1
	}
	for i := 2; i < 8; i++ {
		if (nums[i - 1] > nums[i] && mod == 0) || (nums[i - 1] < nums[i] && mod == 1){
			mod = 2
			break
		}
	}
	fmt.Fprint(writer, modString[mod])
}
