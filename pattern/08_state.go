package main

import (
	"log"
)

func main() {
	doc := Document{"Document 1", &Draft{}}
	draft := Draft{&doc}
	doc.state = &draft

	doc.render()
	doc.publish()
	doc.render()
	doc.publish()
	doc.render()
}

type State interface {
	render()
	publish()
}

type Document struct {
	name  string
	state State
}

func (d *Document) changeState(state State) {
	d.state = state
}

func (d *Document) render() {
	d.state.render()
}

func (d *Document) publish() {
	d.state.publish()
}

type Draft struct {
	document *Document
}

func (d *Draft) render() {
	if true {
		log.Println("Draft:", d.document.name)
		return
	}
	log.Println("Error")
}

func (d *Draft) publish() {
	doc := d.document
	m := Moderation{document: doc}
	d.document.changeState(&m)
}

type Moderation struct {
	document *Document
}

func (m *Moderation) render() {
	if true {
		log.Println("Moderation:", m.document.name)
		return
	}
	log.Println("Error")
}

func (m *Moderation) publish() {
	doc := m.document
	p := Published{document: doc}
	m.document.changeState(&p)
}

type Published struct {
	document *Document
}

func (p *Published) render() {
	if true {
		log.Println("Published:", p.document.name)
		return
	}
	log.Println("Error")
}

func (p *Published) publish() {
}
