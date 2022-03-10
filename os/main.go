package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	filename := pwd + "/test.go"
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(f)
}
