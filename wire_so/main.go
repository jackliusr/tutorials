package main

//https://stackoverflow.com/questions/59568031/why-wire-cant-generate-wire-gen-go/63941411#63941411

func main() {
   /*
    //manual wiring can be used to troubeshooting
     helloSpeaker := NewHelloSpeaker()
	var speaker Speaker =  helloSpeaker
	listener := NewSimpleListener(& speaker)
   */
	listener := InitializeListener()
	listener.WhatIListened()
}

