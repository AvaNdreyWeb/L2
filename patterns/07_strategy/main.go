package main

import "log"

func main() {
	s1 := StrategyA{}
	s2 := StrategyB{}
	ctx := Context{}

	ctx.setStrategy(s1)
	ctx.doSomething()

	ctx.setStrategy(s2)
	ctx.doSomething()
}

type Context struct{ strategy Strategy }

func (c *Context) setStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) doSomething() {
	c.strategy.execute("Some Data")
}

type Strategy interface{ execute(data string) }

type StrategyA struct{}

func (s StrategyA) execute(data string) {
	log.Println("Strategy A", data)
}

type StrategyB struct{}

func (s StrategyB) execute(data string) {
	log.Println("Strategy B", data)
}
