package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
}


func InitDB() {

	db,err = gorm.Open(0
	)
}
