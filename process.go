package main

import (
	"os/exec"
	"bufio"
	"os"
	"log"
)

func NewProcess(command string) *Process {
	return &Process{command: command}
}

type Process struct {
	command string
	cmd     *exec.Cmd
	Pid     int
	Scanner *bufio.Scanner
}

func (p *Process) Run() error {
	cmd := exec.Command("bash", "-c", p.command)
	p.cmd = cmd
	out, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}
	p.Pid = cmd.Process.Pid
	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	p.Scanner = scanner
	cmd.Wait()
	return nil
}

func (p *Process) MustRun() *Process {
	err := p.Run()
	if err != nil {
		log.Fatal("Process.Run:", err)
	}
	return p
}

func (p *Process) Interrupt() error {
	return p.cmd.Process.Signal(os.Interrupt)
}

func (p *Process) Kill() error {
	return p.cmd.Process.Kill()
}

func (p *Process) Wait() {
	p.cmd.Wait()
}
