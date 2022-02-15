package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://119.91.25.133:8080/ping"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("request error %e", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("request error %e", err))
	}
	fmt.Printf("%s", body)
}
