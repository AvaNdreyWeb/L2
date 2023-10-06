package main

import (
	"fmt"
	"strings"
)

type Command interface {
	execute()
}

type Invoker struct {
	cmd Command
}

func (i *Invoker) setCommand(cmd Command) {
	i.cmd = cmd
}

func (i Invoker) executeCommand() {
	i.cmd.execute()
}

type StringsCommand struct {
	reciver StringsReciver
	params  []string
}

func NewStringsCommand(reciver StringsReciver, args ...string) *StringsCommand {
	return &StringsCommand{
		reciver,
		args,
	}
}

func (c StringsCommand) execute() {
	c.reciver.operation(c.params...)
}

type StringsReciver interface {
	operation(args ...string)
}

type ReciverUpper struct{}
type ReciverLower struct{}

func (r ReciverUpper) operation(args ...string) {
	for i := range args {
		args[i] = strings.ToUpper(args[i])
	}
	fmt.Println(strings.Join(args, " "))
}
func (r ReciverLower) operation(args ...string) {
	for i := range args {
		args[i] = strings.ToLower(args[i])
	}
	fmt.Println(strings.Join(args, " "))
}

func main() {
	ru := ReciverUpper{}
	rl := ReciverLower{}
	c1 := NewStringsCommand(rl, "Hello", "World", "Golang")
	c2 := NewStringsCommand(ru, "Hello", "World", "Golang")

	q := CommandQueue{}
	q.Add(c1)
	q.Add(c2)

	i := Invoker{}
	for {
		if cmd, ok := q.Pop(); ok {
			i.setCommand(cmd)
			i.executeCommand()
		} else {
			break
		}
	}
}

type CommandQueue struct {
	queue []Command
}

func (cq *CommandQueue) Add(cmd Command) {
	cq.queue = append(cq.queue, cmd)
}

func (cq *CommandQueue) Pop() (Command, bool) {
	if len(cq.queue) > 0 {
		cmd := cq.queue[0]
		cq.queue = cq.queue[1:]
		return cmd, true
	}
	return StringsCommand{}, false
}
