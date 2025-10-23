package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input) 

    if input == "" {
        fmt.Println(0)
        return
    }

    words := strings.Fields(input)
    fmt.Println(len(words))
}