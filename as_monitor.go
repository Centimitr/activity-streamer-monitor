package main

import "net/http"

var pm PathManager

func init() {
	pm = NewPathManager()
}

func NewASMonitor() *ASMonitor {
	m := &ASMonitor{NewMonitor()}
	m.Init()
	return m
}

type ASMonitor struct {
	*Monitor
}

func (m *ASMonitor) Init() {
	m.Mux.HandleFunc("/as", func(res http.ResponseWriter, req *http.Request) {

	})
}
