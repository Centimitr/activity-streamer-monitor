package main

import (
	"net/http"
	"fmt"
)

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
	mux.HandleFunc("/nodes", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		pid := req.Form.Get("pid")
		operation := req.Form.Get("operation")
		fmt.Println(pid)
		if pid == "" {
			
			return
		}
		switch operation {
		case "start":
		case "interrupt":
		case "terminate":
		}
	})
	m.Mux = mux
}

func (m *Monitor) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), m.Mux)
}
