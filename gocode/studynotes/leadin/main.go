package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("./cat.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	for k, row := range rows {
		if k == 0 || k == 1 {
			continue
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
