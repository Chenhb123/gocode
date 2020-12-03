package main

import (
	"database/sql"
	"fmt"
	"log"

	"datatom.com/ants/logger"
	_ "github.com/lib/pq"
)

func main() {
	db, err := pgConnect("stork", "stork", "192.168.2.80", "odsdb", 14103)
	if err != nil {
		log.Fatal(err)
	}
	sqltxt := fmt.Sprintf("desc extended kt_user")
	rows, err := db.Query(sqltxt)
	if err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}
	for rows.Next() {
		var res1, res2, res3 string
		err = rows.Scan(&res1, &res2, &res3)
		if err != nil {
			fmt.Println("scan")
			log.Fatal(err)
		}
		fmt.Println(res1)
		fmt.Println(res2)
		fmt.Println(res3)
	}
}

func pgConnect(username, password, ip, initDB string, port int) (*sql.DB, error) {
	var db *sql.DB
	pgLink := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username,
		password,
		ip,
		port,
		initDB)
	db, err := sql.Open("postgres", pgLink)
	if err != nil {
		logger.Error.Println(err)
		db.Close()
		return db, err
	}
	err = db.Ping()
	if err != nil {
		logger.Error.Println(err)
		return db, err
	}
	return db, nil
}
