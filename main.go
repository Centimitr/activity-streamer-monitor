package main

import (
	"flag"
	"log"
)

var port = flag.Int("port", 7010, "the port will be used")

func main() {
	m := NewASMonitor()
	log.Println("Listen:", *port)
	m.Listen(*port)
}
