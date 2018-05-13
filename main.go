package main

import (
	"fmt"
	"os"
	"strconv"
)

//var port = flag.Int("port", 7010, "the port will be used")

func main() {
	args := os.Args
	if len(args) >= 2 {
		m := NewASMonitor()
		port, _ := strconv.Atoi(os.Args[1])
		fmt.Println("Listen:", port)
		m.Listen(port)
	}
}
