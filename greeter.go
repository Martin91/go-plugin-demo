package main

import (
	"fmt"
	"os"
	"plugin"
	"plugin_demo/common"
)

type Greeter interface {
	Greet()
	GreetMsg() common.WelcomeMsg
}

func main() {
	lang := "en"

	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	var mod string
	switch lang {
	case "en":
		mod = "./en.so"
	case "ch":
		mod = "./ch.so"
	default:
		fmt.Println("Oh no!")
		os.Exit(1)
	}

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("Oh no, my symbol")
		os.Exit(1)
	}

	greeter.Greet()
	fmt.Printf("Msg is:\n%s", greeter.GreetMsg())
}