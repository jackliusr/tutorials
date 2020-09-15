// key steps during in wire tutorial
// please refer to https://github.com/google/wire/tree/master/_tutorial
package main 

import (
	"fmt"
)

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func main(){
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}


func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start(){
	msg:= e.Greeter.Greet()
	fmt.Println(msg)
}

