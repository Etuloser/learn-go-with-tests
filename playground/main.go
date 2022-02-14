package main

import (
	"fmt"
	"reflect"
)

func hello() []string {
	return nil
}

func main() {

	h := hello
	fmt.Println(reflect.TypeOf(h))
}
