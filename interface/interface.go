package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}
type Dog struct{}
type Cat struct{}
type Llama struct{}

// 类型 Dog 实现了 Animal 接口
func (d Dog) Speak() string {
	return "Woof!"
}
func (c Cat) Speak() string {
	return "Meow!"
}
func (l Llama) Speak() string {
	return "?????"
}

func main() {
	// 一个指针类型可以通过其相关的值类型来访问值类型的方法，但是反过来不行
	animals := []Animal{Dog{}, &Cat{}, Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
