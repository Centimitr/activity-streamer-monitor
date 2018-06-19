package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

var upgrader = websocket.Upgrader{}

func NewMonitor(static string) *Monitor {
	m := &Monitor{}
	m.Init(static)
	return m
}

type Monitor struct {
	Processes map[string]*Process
	Mux       *http.ServeMux
	connm     map[string]*websocket.Conn
}

func (m *Monitor) Init(static string) {
	m.Processes = make(map[string]*Process)
	m.connm = make(map[string]*websocket.Conn)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(static)))
	mux.HandleFunc("/nodes", m.server)
	m.Mux = mux
}

func (m *Monitor) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), m.Mux)
}

func (m *Monitor) server(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	m.connm[r.Host] = c
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func() {
		delete(m.connm, r.Host)
		c.Close()
	}()
	log.Println("Connected:", r.Host)
	for {
		var msg Message
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("Rcve:", err)
			break
		}
		err = c.WriteJSON(m.handle(msg))
		if err != nil {
			log.Println("Snde:", err)
			break
		}
	}

	//for {
	//	mt, req, err := c.ReadMessage()
	//	if err != nil {
	//		log.Println("Rcv:", err)
	//		break
	//	}
	//	var msg Message
	//	json.Unmarshal(req, &msg)
	//	if err != nil {
	//		log.Println("Rcv:", err)
	//		break
	//	}
	//	res, err := json.Marshal(m.handle(msg))
	//	if err != nil {
	//		log.Println("Snd:", err)
	//		break
	//	}
	//	err = c.WriteMessage(mt, res)
	//	if err != nil {
	//		log.Println("Snd:", err)
	//		break
	//	}
	//}
}

func (m *Monitor) Broadcast(msg Message) {
	log.Println("Broadcast:", msg)
	for _, conn := range m.connm {
		conn.WriteJSON(msg)
	}
}

type Message struct {
	Command string   `json:"command"`
	Params  []string `json:"params"`
}

func (m *Monitor) handle(msg Message) (res Message) {
	res.Command = "success"
	switch msg.Command {
	case "start":
		cmd := msg.Params[0]
		p := NewProcess(cmd)
		p.Run()
		go func(p *Process) {
			for p.Scanner.Scan() {
				line := p.Scanner.Text()
				m.Broadcast(Message{
					Command: "stdout",
					Params:  []string{strconv.Itoa(p.Pid), line},
				})
			}
		}(p)
		m.Processes[strconv.Itoa(p.Pid)] = p
	case "interrupt":
		p, ok := m.Processes[msg.Params[0]]
		if ok {
			p.Interrupt()
		}
	case "kill":
		p, ok := m.Processes[msg.Params[0]]
		if ok {
			p.Kill()
		}
	case "remove":
		delete(m.Processes, msg.Params[0])
	}
	return res
}
