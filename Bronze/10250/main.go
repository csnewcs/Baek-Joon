package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var T int

func main() {
	defer writer.Flush()
	
	fmt.Fscan(reader, &T)
	var floor, roomsPerFloor, count int
	for i := 0; i < T; i++ {
		fmt.Fscan(reader, &floor, &roomsPerFloor, &count)
		guestFloor := (count % floor)
		guestRoom := (count / floor) + 1
		if guestFloor == 0 {
			guestFloor = floor
			guestRoom--
		}
		fmt.Fprintf(writer, "%d%02d\n", guestFloor, guestRoom)
	}
}
