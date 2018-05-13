package main

var pm PathManager

func init() {
	pm = NewPathManager()
}

func NewASMonitor() *ASMonitor {
	m := &ASMonitor{NewMonitor("app")}
	m.Init()
	return m
}

type ASMonitor struct {
	*Monitor
}

func (m *ASMonitor) Init() {
}
