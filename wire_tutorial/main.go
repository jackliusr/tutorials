// key steps during in wire tutorial
// please refer to https://github.com/google/wire/tree/master/_tutorial
package main

import (
	"fmt"
        "errors"
        "os"
        "time"
)

type Message string

type Greeter struct {
	Message Message
        Grumpy bool
}

type Event struct {
	Greeter Greeter
}

func main(){
        e, err := InitializeEvent()
        if err != nil {
                fmt.Println("failed to create event: %s\n",err)
                os.Exit(2)
        }
        e.Start()

}

func NewEvent(g Greeter) (Event, error) {
        if g.Grumpy {
                return Event{}, errors.New("could not create event: event greeter is grumpy")
        }
	return Event{Greeter: g}, nil
}
func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
        var grumpy bool
        if time.Now().Unix() % 2 == 0 {
                grumpy = true
        }
	return Greeter{Message: m, Grumpy: grumpy}
}


func (g Greeter) Greet() Message {
        if g.Grumpy {
                return Message("Go away!")
        }
	return g.Message
}

func (e Event) Start(){
	msg:= e.Greeter.Greet()
	fmt.Println(msg)
}
