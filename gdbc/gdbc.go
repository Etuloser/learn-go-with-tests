package gdbc

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Etuloser/ego/pkg/ereflect"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

var MysqlDB *sql.DB
var MysqlErr error

type DatabaseSetting struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

const iniRelativePath = "../my.ini"

func InitDB() {
	wd, _ := os.Getwd()
	iniFilePath := filepath.Join(wd, iniRelativePath)
	cfg, err := ini.Load(iniFilePath)

	if !ereflect.IsNil(err) {
		log.Fatal(err.Error())
	}

	dbSetting := new(DatabaseSetting)
	cfg.Section("mysql").MapTo(dbSetting)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.Port,
		dbSetting.Database,
		dbSetting.Charset)
	MysqlDB, MysqlErr = sql.Open("mysql", dsn)
	if !ereflect.IsNil(MysqlErr) {
		log.Println("dsn: " + dsn)
		panic("数据源配制不正确: " + MysqlErr.Error())
	}
	// 最大连接数
	MysqlDB.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDB.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDB.SetConnMaxLifetime(100 * time.Second)

	if MysqlErr = MysqlDB.Ping(); nil != MysqlErr {
		panic("数据库链接失败: " + MysqlErr.Error())
	}
}

func RawQuery() {
	rows, _ := MysqlDB.Query("select * from subents")
	if rows == nil {
		return
	}
	id := 0
	subnet := ""
	fmt.Println(rows)
	fmt.Println(rows)
	for rows.Next() {
		rows.Scan(&id, &subnet)
		fmt.Println(id, subnet)
	}
}
