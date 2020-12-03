package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// /opt/danastudio/ants/tdc_simple_load/ods_tdcmysql0320_1584666135__2594ab33_4564_473c_a127_c737ef3ad81a
// curl -C - -X PUT -T  "http://192.168.90.91:31573/webhdfs/v1/tmp/detuo/ods_tdcmysql0320_1584666135?op=CREATE&data=true&permission=777&guardian_access_token=cKjIZGNM8jGgqQrtss4f-WHVTFZ0.TDH" -H "Content-Type:application/octet-stream"

// curl -C - -X PUT -T $normalpath "http://192.168.90.91:31573/webhdfs/v1/tmp/detuo/ods_tdcmysql0320_1584666135?op=CREATE&data=true&permission=777&guardian_access_token=cKjIZGNM8jGgqQrtss4f-WHVTFZ0.TDH" -H "Content-Type:application/octet-stream"
// python /opt/dana/datax/bin/datax.py /var/dana/dodox/filemanager/file/danastudio-unsubmit/hive/admin_user_落地_tdcmysql0320fdaT.datax/admin_user_落地_tdcmysql0320fdaT.datax

// jdbc:hive2://dn8029:2181/;serviceDiscoveryMode=zooKeeper;zooKeeperNamespace=hiveserver2
// ddp,1.2、3.1、tdh、tdc、cdh、华为Fs
func main() {

	url := "http://192.168.2.80:21613/pool/hive/executeSql"
	method := "POST"

	payload := strings.NewReader("jdbcUrl=jdbc:hive2://dn8029:2181/;serviceDiscoveryMode=zooKeeper;zooKeeperNamespace=hiveserver2&username=hive&password=123456&sql=show databases")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
