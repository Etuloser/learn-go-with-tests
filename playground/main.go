package main

import (
	"fmt"
	"learn-go-with-tests/runtime"
)

func main() {
	fileName := runtime.GetFileName()
	fmt.Println(fileName)
}
