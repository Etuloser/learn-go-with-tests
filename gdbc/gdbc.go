package gdbc

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

func InitDB(iniFilePath string) {
	cfg, err := ini.Load(iniFilePath)

	if err != nil {
		panic(err)
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
	if MysqlErr != nil {
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

func RawQuery(rawSql string) {
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
