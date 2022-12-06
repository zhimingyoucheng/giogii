package main

import (
	"flag"
	"giogii/src/check"
	"giogii/src/lock"
	"strings"
	"testing"
)

func TestM(t *testing.T) {
	var sourceUserInfo string
	var sourceSocket string
	var targetUserInfo string
	var targetSocket string
	var parameter string
	var bigTrx string

	/*flag.StringVar(&sourceUserInfo, "s", "root:drACgwoqtM", "")
	flag.StringVar(&sourceSocket, "si", "172.17.128.49:13336", "")
	flag.StringVar(&targetUserInfo, "t", "wjy_root:Wjy123456", "")
	flag.StringVar(&targetSocket, "ti", "rm-2ze5j9oqx3x70jzd94o.mysql.rds.aliyuncs.com:3306", "")
	flag.StringVar(&parameter, "c", "8c32gb", "")*/

	flag.StringVar(&sourceUserInfo, "s", "admin:!QAZ2wsx", "")
	flag.StringVar(&sourceSocket, "si", "172.17.139.27:16310", "")
	flag.StringVar(&targetUserInfo, "t", "root:!QAZ2wsx", "")
	flag.StringVar(&targetSocket, "ti", "172.17.140.3:16310", "")
	flag.StringVar(&parameter, "c", "", "")
	flag.StringVar(&bigTrx, "m", "m", "")

	flag.Parse()

	if strings.Trim(parameter, " ") == "c" {
		check.InitCheckParameterConf(sourceUserInfo, sourceSocket, "greatrds", targetUserInfo, targetSocket, "information_schema")
		check.DoCheckParameter(parameter)
	} else if strings.Trim(bigTrx, " ") == "m" {
		lock.InitConf(sourceUserInfo, sourceSocket, "performance_schema")
		lock.DoMonitorLock()
	} else {
		check.InitCheckConsistentConf(sourceUserInfo, sourceSocket, "information_schema", targetUserInfo, targetSocket, "information_schema")
		check.DoCheck()
	}
}
