package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	tt := time.Now()
	var db *sql.DB
	pgLink := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"stork",
		"stork",
		"192.168.90.48",
		14103,
		"yyqqx_storkdb")

	db, err := sql.Open("postgres", pgLink)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("连接耗时：", time.Since(tt))
	tt = time.Now()
	sqltxt := "create table public.test724(id int, name text)"
	// sqltxt := ""
	_, err = db.Exec(sqltxt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("执行耗时：", time.Since(tt))
	// fmt.Println("执行耗时：", time.Since(tt))
	// sqltxt := fmt.Sprintf(`
	// SELECT u.datname  FROM pg_catalog.pg_database u where u.datname='demo_dmdb'`)
	// sqltxt := "SELECT datname FROM pg_database"
	// rows, err := db.Query(sqltxt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// var field string
	// for rows.Next() {
	// 	rows.Scan(&field)
	// 	fmt.Println(field)
	// }
}
