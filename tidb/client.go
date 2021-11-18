package tidb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Client struct {
	db *sql.DB
}

func NewClient() Client {
	db, err := sql.Open("mysql", "root@tcp(172.16.6.173:31010)/cmtest")

	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(200)

	if err := db.Ping(); err != nil {
		log.Panicf("open database fail: %v", err)
	}
	fmt.Println("connect success")
	return Client{
		db,
	}
}
