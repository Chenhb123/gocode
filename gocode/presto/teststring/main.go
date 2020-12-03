package main

import (
	"fmt"
	"strings"
)

// Person 。
type Person struct {
	Name string
	Address
}

// Address 。
type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

// String .
func (a *Address) String() string {
	return a.Number + " " + a.Street + "\n" + a.City + ", " + a.State + " " + a.Zip + "\n"
}

// String .
func (p *Person) String() string {
	return p.Name
}

func main() {
	result := `1139366  /inceptor1/user/hive/warehouse/odsdb.db/hive/carinfo`
	resArr := strings.Split(result, "\n")
	resArr = strings.Split(resArr[len(resArr)-1], "\t")
	fmt.Println("resArr:", resArr)
}
