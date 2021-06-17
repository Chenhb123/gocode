package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type medal struct {
	name   string
	gold   int
	silver int
	bronze int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	medals := make(map[string]medal)
	var name string
	var gold, silver, bronze int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= n; i++ {
		line, _, _ := reader.ReadLine()
		lFields := strings.Fields(string(line))
		if len(lFields) != 4 {
			continue
		}
		name = lFields[0]
		gold, _ = strconv.Atoi(lFields[1])
		silver, _ = strconv.Atoi(lFields[2])
		bronze, _ = strconv.Atoi(lFields[3])
		var medal medal
		medal.name = name
		medal.gold = gold
		medal.silver = silver
		medal.bronze = bronze
		medals[name] = medal
	}
	result := sortMedal(medals)

	for _, v := range result {
		fmt.Println(v)
	}
}

func sortMedal(medals map[string]medal) []string {
	var result []string
	// 排序规则--优先级：gold > silver > bronze > name
	// 先实现最复杂
	medMap := make(map[string]bool)
	for {
		if len(result) == len(medals) {
			break
		}
		var max medal
		for i := range medals {
			if _, ok := medMap[medals[i].name]; ok {
				continue
			}
			max = medals[i]
			for m := range medals {
				if _, ok := medMap[medals[m].name]; ok {
					continue
				}
				if medals[m].gold > max.gold {
					max = medals[m]
					continue
				}
				if medals[m].gold == max.gold && medals[m].silver > max.silver {
					max = medals[m]
					continue
				}
				if medals[m].gold == max.gold && medals[m].silver == max.silver && medals[m].bronze > max.bronze {
					max = medals[m]
					continue
				}
				if medals[m].gold == max.gold && medals[m].silver == max.silver && medals[m].bronze == max.bronze {
					// 判定字典顺序
					var temp []string
					temp = append(temp, max.name, medals[m].name)
					sort.Strings(temp)
					if temp[0] == medals[m].name {
						max = medals[m]
					}
				}
			}
			medMap[max.name] = true
			result = append(result, max.name)
		}
	}
	return result
}
