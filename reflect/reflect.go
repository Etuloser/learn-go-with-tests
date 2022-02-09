package main

import (
	"fmt"
	"reflect"
)

var a int

type Dog struct{}

var Animal interface{}

var intPtr *int

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
	var i int = 10
	intPtr = &i
	v := reflect.ValueOf(intPtr)
	k := v.Kind()
	fmt.Println(v, k)
}
