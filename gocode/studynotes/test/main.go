package main

import (
	"log"
	"strings"

	"datatom.com/ants/common"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var err error
	f := excelize.NewFile()
	styleCat, _ := f.NewStyle(`{
		"border": [{"type":"top","color":"000000","style":2},
			{"type":"bottom","color":"000000","style":2},
			{"type":"right","color":"000000","style":2}]
		}`)
	styleTop, _ := f.NewStyle(`{
		"border": [{"type":"top","color":"000000","style":2},
			{"type":"bottom","color":"000000","style":2},
			{"type":"right","color":"000000","style":5}]
		}`)
	styleText, _ := f.NewStyle(`{"alignment": {"horizontal": "centerContinuous"}}`)
	for k, v := range common.MetaTopCat {
		err = f.SetCellValue("Sheet1", k, v)
		if err != nil {
			log.Fatal(err)
		}

	}
	for k, v := range common.MetaCategories {
		err = f.SetCellValue("Sheet1", k, v)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(k, "M") ||
			strings.Contains(k, "P") ||
			strings.Contains(k, "T") {
			f.SetCellStyle("Sheet1", k, k, styleTop)
			continue
		}
		f.SetCellStyle("Sheet1", k, k, styleCat)
	}
	for k, v := range common.MetaMergeCell1 {
		f.MergeCell("Sheet1", k, v)
		f.SetCellStyle("Sheet1", v, v, styleTop)
		f.SetCellStyle("Sheet1", k, k, styleText)
	}
	index := f.NewSheet("Sheet2")
	for k, v := range common.MetaMergeCell2 {
		f.SetCellValue("Sheet2", k, v)
		f.SetCellStyle("Sheet2", k, k, styleCat)
	}
	f.SetActiveSheet(index)
	err = f.SaveAs("./cat.xlsx")
	if err != nil {
		log.Fatal(err)
	}

}
