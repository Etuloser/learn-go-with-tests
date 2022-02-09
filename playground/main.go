package main

import "fmt"

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type tHelper interface {
	Helper()
}

var t TestingT

func main() {
	var a interface{}
	value, ok := a.(string)
	if !ok {
		fmt.Println("It's not k for type string")
		return
	}
	fmt.Println("This value is ", value)
}
