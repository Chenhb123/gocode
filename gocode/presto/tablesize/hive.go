package main

//功能 ： 验证获取表条数占用内存情况
import (
	"fmt"
	hive "github.com/dazheng/gohive"
	"log"
	"strings"
)

func main() {
	IP := "192.168.80.20"
	Port := 10000
	Username := "hive"
	Password := "hive"
	database := "odsdb"
	//table := "test_jj"
	addStr := fmt.Sprintf("%s:%d", IP, Port)
	db, err := hive.ConnectWithUser(addStr, Username, Password, hive.DefaultOptions)
	if err != nil {
		fmt.Println("---1---err=%s-----", err.Error())
	}
	defer db.Close()
	sqltxt := fmt.Sprintf("use %s", database)
	_, err = db.Exec(sqltxt)
	if err != nil {
		fmt.Println("--2----err=%s-----", err.Error())
	}
	sqltxt = "show tables"
	rows, err := db.Query(sqltxt)
	if err != nil {
		log.Fatal(err)
	}
	var tables []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		tables = append(tables, table)
	}
	var numRows string
	for i := 0; i < 100000000; i++ {
		// sqltxt1 := fmt.Sprintf("desc formatted %s", table)
		// rows, err := db.Query(sqltxt1)
		// if err != nil {
		// 	fmt.Println("---3---err=%s-----", err.Error())
		// }
		// for rows.Next() {
		// 	var name, aa, tt string
		// 	err = rows.Scan(&name, &aa, &tt)
		// 	if err != nil {
		// 		fmt.Println("---4---err=%s-----", err.Error())
		// 	}
		// 	if !strings.HasPrefix(aa, "numRows") {
		// 		continue
		// 	}
		// 	fmt.Println("------tt=%v-----", tt)
		// }
		for _, v := range tables {
			sqltxt := fmt.Sprintf("desc formatted `%s`", v)
			rows, err := db.Query(sqltxt)
			if err != nil {
				fmt.Println("---3---err=%s-----", err.Error())
				continue
			}
			for rows.Next() {
				var name, aa, tt string
				err = rows.Scan(&name, &aa, &tt)
				if err != nil {
					fmt.Println("---4---err=%s-----", err.Error())
				}
				if !strings.HasPrefix(aa, "numRows") {
					// numRows = aa
					continue
				}
				numRows = tt
				break
				// fmt.Println("------tt=%v-----", tt)
			}
			if numRows == "0" {
				sqltxt = fmt.Sprintf("ANALYZE TABLE  `%s` COMPUTE STATISTICS", v)
				_, err := db.Exec(sqltxt)
				if err != nil {
					fmt.Println("---3---err=%s-----", err.Error())
				}
				sqltxt := fmt.Sprintf("desc formatted `%s`", v)
				rows, err = db.Query(sqltxt)
				if err != nil {
					fmt.Println("---3---err=%s-----", err.Error())
					continue
				}
				for rows.Next() {
					var name, aa, tt string
					err = rows.Scan(&name, &aa, &tt)
					if err != nil {
						fmt.Println("---4---err=%s-----", err.Error())
					}
					if !strings.HasPrefix(aa, "numRows") {
						// numRows = aa
						continue
					}
					numRows = tt
					break
					// fmt.Println("------tt=%v-----", tt)
				}
			}
			// fmt.Println("numRows:", numRows)
			fmt.Println("i:", i, "tablename:", v, "numRows:", numRows)
		}
	}
	fmt.Println("------结束-----")

}
