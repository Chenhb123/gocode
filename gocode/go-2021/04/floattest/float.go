package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	test := []string{"all_types",
		"all_types_copy1",
		"all_types_copy2",
		"all_types_copy3",
		"all_types_copy4",
		"all_types_copy5"}
	bytes, err := json.Marshal(test)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

	str := []string{"varchartime", "view_accident", "view_application", "view_carinfo",
		"view_cube", "view_czyperson", "view_key", "view_personinfo", "view_select", "view_stringtime", "view_student", "view_time",
		"view_xllperson", "wj", "xiaoyi1", "xiaoyi1_copy2", "xiaoyi初始表不要动复制再用", "xll1", "xll_person", "yyq_gx_gaussdb_mysql_kbts",
		"yyq_mysql_kbts1", "yyq_pl_change1", "yyq_pl_change2", "yyq_ss_application_and_alltypes1", "yyq_ss_application_and_alltypes2",
		"yyq_ss_application_and_alltypes3", "yyq_ss_application_and_alltypes4", "yyq_ss_application_and_alltypes5",
		"yyq_ss_application_and_alltypes6", "yyq_ss_application_and_alltypes7", "yzsxx", "zhuanti", "zhuanti425", "zhujian", "zl3", "zltestmysql", "zt1", "zt12zl", "ztsc0209", "zy8831test", "zyappyyyymmdd", "zymysqlpk", "zymysqlpku1", "zymysqlpku2", "zymysqlunique", "zynotapply", "zyperson100", "zystom_zltime_add2", "zytest", "zytest471", "zytest471mod", "zytestbool", "zytestjb", "zytestview471", "zyttom_zlstring_add", "zyzl_stork_time"}
	fmt.Println(len(str))
}
