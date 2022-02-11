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
type Subnet struct {
	Id     int64  `db:"id"`
	Subent string `db:"网段"`
	Mask   string `db:"掩码"`
}

func InitDB(iniFilePath string) {
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
func setup() {
	wd, _ := os.Getwd()
	iniFilePath := filepath.Join(wd, "../my.ini")
	InitDB(iniFilePath)
}

func ExampleSearch() {
	setup()
	subnet := new(Subnet)
	rows := MysqlDB.QueryRow("select id,subnet,mask from subnets")
	if err := rows.Scan(&subnet.Id, &subnet.Subent, &subnet.Mask); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(subnet.Id, subnet.Subent, subnet.Mask)
}
func ExampleAdd() {
	setup()
	ret, err := MysqlDB.Exec("insert into subnets(subnet,mask) value(?,?)", "12556", "16")
	if err != nil {
		log.Fatal(err.Error())
	}

	//插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	fmt.Println("LastInsertID:", lastInsertID)

	//影响行数
	rowsaffected, _ := ret.RowsAffected()
	fmt.Println("RowsAffected:", rowsaffected)
}
func ExampleUpdate() {

	ret, _ := MysqlDB.Exec("UPDATE subnets set subnet=? where id=?", "100", 1)
	upd_nums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", upd_nums)
}

func ExampleDel() {

	ret, _ := MysqlDB.Exec("delete from subnets where id=?", 1)
	del_nums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", del_nums)
}
