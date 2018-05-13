package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) >= 2 {
		m := NewASMonitor()
		port, _ := strconv.Atoi(os.Args[1])
		fmt.Println("Listen:", port)
		m.Listen(port)
	}
}
