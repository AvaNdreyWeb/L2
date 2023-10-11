package main

import "log"

const (
	StatusOK byte = iota
	StatusDisc
	StatusErr
	StatusRecovered
)

func main() {
	cl1 := Client{Node{StatusErr, "Client 1"}}
	cl2 := Client{Node{StatusDisc, "Client 2"}}
	dc1 := DataCenter{Node{StatusOK, "Data Center 1"}}
	dc2 := DataCenter{Node{StatusErr, "Data Center 2"}}
	core := Core{Node{StatusOK, "Core"}}
	net := []Element{cl1, cl2, dc1, dc2, core}

	sv := StatusVisitor{}

	for _, e := range net {
		e.accept(sv)
	}
}

type Visitor interface {
	visitCore(c Core)
	visitClient(cl Client)
	visitDataCenter(dc DataCenter)
}

type Element interface {
	accept(v Visitor)
}

type StatusVisitor struct{}

func (v StatusVisitor) visitCore(c Core) {
	v.logStatus(c.status, c.name)
}

func (v StatusVisitor) visitClient(cl Client) {
	v.logStatus(cl.status, cl.name)
}

func (v StatusVisitor) visitDataCenter(dc DataCenter) {
	if dc.status != StatusOK {
		dc.restart()
	}
	v.logStatus(dc.status, dc.name)
}

func (v StatusVisitor) logStatus(status byte, name string) {
	switch status {
	case StatusOK:
		log.Printf(" INFO: %-20s [OK]", name)
	case StatusDisc:
		log.Printf(" INFO: %-20s [DISCONNECTED]", name)
	case StatusErr:
		log.Printf("ERROR: %-20s [ERROR]", name)
	case StatusRecovered:
		log.Printf(" INFO: %-20s [RECOVERED]", name)
	}
}

type Node struct {
	status byte
	name   string
}

type Core struct{ Node }
type Client struct{ Node }
type DataCenter struct{ Node }

func (c Core) accept(v Visitor) {
	v.visitCore(c)
}

func (cl Client) accept(v Visitor) {
	v.visitClient(cl)
}

func (dc DataCenter) accept(v Visitor) {
	v.visitDataCenter(dc)
}

func (dc *DataCenter) restart() {
	dc.status = StatusRecovered
}
