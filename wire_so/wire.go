// +build wireinject
package main
import (
	"github.com/google/wire"
	"fmt"
)

type Speaker interface {
	Say()
}

type HelloSpeaker struct {
	word string
}

func NewHelloSpeaker() *HelloSpeaker {
	return &HelloSpeaker{
		word: "Hello, World!!",
	}
}

func (s *HelloSpeaker) Say() {
	fmt.Printf("%s\n", s.word)
}

type Listener interface {
	WhatIListened()
}

type SimpleListener struct {
	speaker *Speaker
}

func NewSimpleListener(speaker Speaker) *SimpleListener {
	return &SimpleListener{
		speaker: &speaker,
	}
}

func (l *SimpleListener) WhatIListened() {
	(*l.speaker).Say()
}


func InitializeListener() Listener {
	wire.Build(
		NewHelloSpeaker, //*HelloSpeaker, HelloSpeaker doesn't implement Speaker interfaces, conversion is needed in next step
		wire.Bind(new(Speaker),new(*HelloSpeaker)), //this binding can provide Speaker, not *Speaker
		NewSimpleListener, //scenario as NewHelloSpeaker
		//provide Listener, an extra method is needed if *Listener is wanted.
		wire.Bind(new(Listener), new(*SimpleListener)),
	)
	return nil
}
