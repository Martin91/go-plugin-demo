package main

import (
	"fmt"
	"plugin_demo/common"
)

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Zeqin")
}

func (g greeting) GreetMsg() common.WelcomeMsg {
	return "msg: Hello, my friend"
}

var Greeter greeting