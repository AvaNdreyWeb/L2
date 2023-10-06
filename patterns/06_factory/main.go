package main

import (
	"log"
)

type Creator interface {
	someOperation()
	createProduct() Product
}

type ProductACreator struct{ Creator }

func (pac ProductACreator) someOperation() {
	p := ProductA{}
	p.doStaff()
}

func (pac ProductACreator) createProduct() ProductA {
	log.Println("Created new Product A")
	return ProductA{}
}

type ProductBCreator struct{ Creator }

type Product interface {
	doStaff()
}

func (pbc ProductBCreator) someOperation() {
	p := ProductB{}
	p.doStaff()
}

func (pbc ProductBCreator) createProduct() ProductB {
	log.Println("Created new Product B")
	return ProductB{}
}

type ProductA struct{}

func (p ProductA) doStaff() {
	log.Println("Product A doing staff!")
}

type ProductB struct{}

func (p ProductB) doStaff() {
	log.Println("Product B doing staff!")
}

func main() {
	pac := ProductACreator{}
	pbc := ProductBCreator{}

	pac.someOperation()
	pbc.someOperation()

	pa := pac.createProduct()
	pb := pbc.createProduct()

	pa.doStaff()
	pb.doStaff()
}
