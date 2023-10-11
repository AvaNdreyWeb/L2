package main

import "fmt"

func main() {
	pb := PizzaBuilder{}
	pd := NewPizzaDirector(pb)

	margaritaM := pd.makeMargarita(24)
	fmt.Println("Margarita M:")
	margaritaM.Info()

	chickenBBQXL := pd.makeChickenBBQ(32)
	fmt.Println()
	fmt.Println("ChickenBBQ XL:")
	chickenBBQXL.Info()
}

type Pizza struct {
	diameter   int
	peperoni   bool
	chicken    bool
	bacon      bool
	ham        bool
	chorizo    bool
	cheddar    bool
	mozarella  bool
	chili      bool
	tomatoes   bool
	pineapples bool
	shrooms    bool
	bbq        bool
	ranch      bool
}

func (p Pizza) Info() {
	fmt.Println("diameter:  ", p.diameter)
	fmt.Println("peperoni:  ", p.peperoni)
	fmt.Println("chicken:   ", p.chicken)
	fmt.Println("bacon:     ", p.bacon)
	fmt.Println("ham:       ", p.ham)
	fmt.Println("chorizo:   ", p.chorizo)
	fmt.Println("cheddar:   ", p.cheddar)
	fmt.Println("mozarella: ", p.mozarella)
	fmt.Println("chili:     ", p.chili)
	fmt.Println("tomatoes:  ", p.tomatoes)
	fmt.Println("pineapples:", p.pineapples)
	fmt.Println("shrooms:   ", p.shrooms)
	fmt.Println("bbq:       ", p.bbq)
	fmt.Println("ranch:     ", p.ranch)
}

type PizzaDirector struct {
	b PizzaBuilder
}

func NewPizzaDirector(pb PizzaBuilder) PizzaDirector {
	return PizzaDirector{pb}
}

func (d PizzaDirector) makeMargarita(diameter int) Pizza {
	d.b.Reset()
	d.b.SetDiameter(diameter)
	d.b.AddMozarella()
	d.b.AddTomatoes()
	return d.b.GetPizza()
}

func (d PizzaDirector) makeChickenBBQ(diameter int) Pizza {
	d.b.Reset()
	d.b.SetDiameter(diameter)
	d.b.AddBBQ()
	d.b.AddChicken()
	d.b.AddChili()
	d.b.AddTomatoes()
	d.b.AddShrooms()
	d.b.AddMozarella()
	return d.b.GetPizza()
}

type PizzaBuilder struct {
	Pizza
}

func (b *PizzaBuilder) SetDiameter(d int) {
	b.diameter = d
}

func (b *PizzaBuilder) AddPeperoni() {
	b.peperoni = true
}

func (b *PizzaBuilder) AddChicken() {
	b.chicken = true
}

func (b *PizzaBuilder) AddBacon() {
	b.bacon = true
}

func (b *PizzaBuilder) AddHam() {
	b.ham = true
}

func (b *PizzaBuilder) AddChorizo() {
	b.chorizo = true
}

func (b *PizzaBuilder) AddCheddar() {
	b.cheddar = true
}

func (b *PizzaBuilder) AddMozarella() {
	b.mozarella = true
}

func (b *PizzaBuilder) AddChili() {
	b.chili = true
}

func (b *PizzaBuilder) AddTomatoes() {
	b.tomatoes = true
}

func (b *PizzaBuilder) AddPineapples() {
	b.pineapples = true
}

func (b *PizzaBuilder) AddShrooms() {
	b.shrooms = true
}

func (b *PizzaBuilder) AddBBQ() {
	b.bbq = true
}

func (b *PizzaBuilder) AddRanch() {
	b.ranch = true
}

func (b *PizzaBuilder) Reset() {
	b.diameter = 0
	b.peperoni = false
	b.chicken = false
	b.bacon = false
	b.ham = false
	b.chorizo = false
	b.cheddar = false
	b.mozarella = false
	b.chili = false
	b.tomatoes = false
	b.pineapples = false
	b.shrooms = false
	b.bbq = false
	b.ranch = false
}

func (b PizzaBuilder) GetPizza() Pizza {
	pizza := Pizza{
		diameter:   b.diameter,
		peperoni:   b.peperoni,
		chicken:    b.chicken,
		bacon:      b.bacon,
		ham:        b.ham,
		chorizo:    b.chorizo,
		cheddar:    b.cheddar,
		mozarella:  b.mozarella,
		chili:      b.chili,
		tomatoes:   b.tomatoes,
		pineapples: b.pineapples,
		shrooms:    b.shrooms,
		bbq:        b.bbq,
		ranch:      b.ranch,
	}
	return pizza
}
