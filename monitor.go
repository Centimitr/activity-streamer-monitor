package main

import (
	"net/http"
	"fmt"
)

func NewMonitor() *Monitor {
	m := &Monitor{}
	m.Init()
	return m
}

type Monitor struct {
	Processes []*Process
	Mux       *http.ServeMux
}

func (m *Monitor) Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

	})
	m.Mux = mux
}

func (m *Monitor) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), m.Mux)
}
