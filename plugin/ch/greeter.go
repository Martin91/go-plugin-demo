package main

import "fmt"
import "plugin_demo/common"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好世界")
}

func (g greeting) GreetMsg() common.WelcomeMsg {
	return "你好，这是我的消息"
}

var Greeter greeting