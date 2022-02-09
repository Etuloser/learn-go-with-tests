package main

import (
	"fmt"
	"reflect"
)

var a int

type Dog struct{}

var Animal interface{}

func main() {
	typeOfA := reflect.TypeOf(a)
	typeOfDog := reflect.TypeOf(Dog{})
	typeOfAnimal := reflect.TypeOf(Dog{})

	fmt.Println(typeOfA.Name())
	fmt.Println(typeOfDog.Name())
	fmt.Println(typeOfAnimal.Kind())

	dog := &Dog{}
	Animal = dog

	fmt.Println(reflect.TypeOf(Animal).Kind())

}
