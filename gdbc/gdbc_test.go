package gdbc

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {
	wd, _ := os.Getwd()
	iniFilePath := filepath.Join(wd, "../my.ini")
	InitDB(iniFilePath)
}

func TestExample(t *testing.T) {
	ExampleSearch()
	ExampleAdd()
	ExampleUpdate()
	ExampleDel()
}
