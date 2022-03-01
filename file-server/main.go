package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(p)
	http.Handle("/", http.FileServer(http.Dir(p)))
	err := http.ListenAndServe(":10486", nil)
	if err != nil {
		fmt.Println(err)
	}
}
