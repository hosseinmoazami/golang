package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	display.Display("Hello")
	msg.Hi("Hi from main")
	msg.Exciting("Exciting")
}
