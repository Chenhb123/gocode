package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	datas := []map[string]interface{}{
		{"bigint": 1, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港1"},
		{"bigint": 2, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港2"},
		{"bigint": 3, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港3"},
		{"bigint": 4, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港4"},
		{"bigint": 5, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港5"},
		{"bigint": 6, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港6"},
		{"bigint": 7, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港7"},
		{"bigint": 8, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港8"},
		{"bigint": 9, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港9"},
		{"bigint": 10, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港10"},
		{"bigint": 11, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港11"},
		{"bigint": 12, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港12"},
		{"bigint": 13, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港13"},
		{"bigint": 14, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港14"},
		{"bigint": 15, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港15"},
		{"bigint": 16, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港16"},
		{"bigint": 17, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港17"},
		{"bigint": 18, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港18"},
		{"bigint": 19, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港19"},
		{"bigint": 20, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港20"},
		{"bigint": 21, "abit": true, "bool": 1, "boolean": 1, "varchar_1": "香港21"},
	}
	// 获取分隔文件数
	files := len(datas) / 5
	if len(datas)%5 > 0 {
		files += 1
	}
	// 开始分隔处理
	var wg sync.WaitGroup
	var fileSize [][]map[string]interface{}
	for len(datas) > 0 {
		if len(datas) <= 5 {
			fileSize = append(fileSize, datas[:])
			break
		}
		fileSize = append(fileSize, datas[:5])
		datas = datas[5:]
	}
	for _, v := range fileSize {
		fmt.Println(v)
	}
	for i := range fileSize {
		wg.Add(1)
		//item := datas[5*i:5*(i+1)] // 取数据
		item := fileSize[i]
		go func(index int, item []map[string]interface{}) {
			DwsLoadDowdExcel(index, item)
			wg.Done()
		}(i, item) // 分别处理每5条的数据
	}
	wg.Wait()
}

//下载离线数据为excel格式的
func DwsLoadDowdExcel(name int, datas []map[string]interface{}) {
	//var err error
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	beginTimestamp := time.Now()
	var header []string
	for i := range datas[0] {
		header = append(header, i)
	}
	sort.Strings(header) //给切片排序
	for lineNum, v := range datas {
		clumnNum := 0
		for i := 0; i < len(v); i++ {
			clumnNum++
			sheetPosition := Divs(clumnNum) + strconv.Itoa(lineNum+2)
			//写入数据
			f.SetCellValue("Sheet1", sheetPosition, v[header[i]])
			sheetPosit := Divs(clumnNum) + "1"
			//设置字段
			f.SetCellValue("Sheet1", sheetPosit, header[i])
		}
		f.SetActiveSheet(index)
		//以10行为分割点，名称规则(book1.xlsx ,book2.xlsx,book3.xlsx.....)
		if err := f.SaveAs(fmt.Sprintf("book%d.xlsx", name)); err != nil {
			fmt.Println(err)
		}
	}
	endTimestamp := time.Now()
	elapsed := endTimestamp.Sub(beginTimestamp)
	fmt.Printf("Go: Writing  cells of data takes %v seconds", elapsed)
}

// Div 数字转字母
func Divs(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
