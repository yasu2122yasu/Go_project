package service

import (
	"database/sql"
	"fmt"
	"go-iris-sample/model"
	"log"
	"os"

	"github.com/go-gorp/gorp"

	_ "github.com/go-sql-driver/mysql"
)

// gorp初期化処理
func InitDb() *gorp.DbMap {
	dbUserName := "DB_USERNAME"
	dbPassWord := "DB_PASSWORD"
	dbHost := "DB_HOST"
	dbPort := "DB_PORT"
	dbDbName := "DB_DBNAME"
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbUserName, dbPassWord, dbHost, dbPort, dbDbName))
	if err != nil {
		fmt.Printf("error! can't open sql driver %v", err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// テーブル作成
	dbmap.AddTableWithName(model.User{}, "user").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		fmt.Printf("error! can't craete table: %v", err)
	}

	// ログの取得
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-iris-sample:", log.Lmicroseconds))

	return dbmap
}
