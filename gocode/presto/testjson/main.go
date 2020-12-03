package main

import (
	"encoding/json"
	"fmt"
	"log"

	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	res := `{"0": "{\"database_name\":{\"0\":\"default\",\"1\":\"discover\",\"2\":\"system\"}}"}`
	js, err := simplejson.NewJson([]byte(res))
	if err != nil {
		log.Fatal(err)
	}
	rec1, err := js.Get("0").String()
	if err != nil {
		log.Fatal(err)
	}
	rec2 := make(map[string]map[string]string, 0)
	err = json.Unmarshal([]byte(rec1), &rec2)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range rec2 {
		for _, r := range v {
			fmt.Println(r)
		}
	}
}
