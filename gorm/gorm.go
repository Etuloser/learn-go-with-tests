package gorm

import (
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type DatabaseSetting struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

func setup() {
	wd, _ := os.Getwd()
	iniFilePath := filepath.Join(wd, "../my.ini")
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		panic(err)
	}
	databaseSetting := new(DatabaseSetting)
	cfg.Section("mysql").MapTo(databaseSetting)
}

func InitDB() {
	setup()
}
