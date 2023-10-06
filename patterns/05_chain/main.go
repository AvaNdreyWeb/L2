package main

import "log"

type HandlerAuth struct{ BaseHandler }

func (h HandlerAuth) handle(r Request) {
	if !(r.username == "valid" && r.password == "valid") {
		log.Println("ERROR: Wrong username or password")
		return
	}
	h.next.handle(r)
}

type HandlerAdmin struct{ BaseHandler }

func (h HandlerAdmin) handle(r Request) {
	if !r.isAdmin {
		log.Println("ERROR: Permisson denied")
		return
	}
	h.next.handle(r)
}

func main() {
	bh := BaseHandler{nil}
	h1 := HandlerAuth{}
	h2 := HandlerAdmin{}

	h1.setNext(&h2)
	h2.setNext(&bh)

	validUserRequest := Request{
		"valid",
		"valid",
		false,
	}
	invalidAdminRequest := Request{
		"invalid",
		"invalid",
		true,
	}
	validAdminRequest := Request{
		"valid",
		"valid",
		true,
	}
	h1.handle(validUserRequest)
	h1.handle(invalidAdminRequest)
	h1.handle(validAdminRequest)
}

type Request struct {
	username string
	password string
	isAdmin  bool
}

type Handler interface {
	setNext(h Handler)
	handle(r Request)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) setNext(h Handler) {
	b.next = h
}

func (b *BaseHandler) handle(r Request) {
	if b.next != nil {
		b.next.handle(r)
		return
	}
	log.Println(" INFO: Done!")
}
