package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{}

func NewMonitor(static string) *Monitor {
	m := &Monitor{}
	m.Init(static)
	return m
}

type Monitor struct {
	Processes []*Process
	Mux       *http.ServeMux
}

func (m *Monitor) Init(static string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(static)))
	mux.HandleFunc("/nodes", m.server)
	m.Mux = mux
}

func (m *Monitor) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), m.Mux)
}

func (m *Monitor) server(w http.ResponseWriter, r *http.Request) {
	log.Println("will upgrade!")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("upgraded!")
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %d %s", mt, message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
