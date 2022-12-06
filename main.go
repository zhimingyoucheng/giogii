package main

import (
	"flag"
	"giogii/src/check"
	"giogii/src/lock"
	"strings"
)

func main() {
	var sourceUserInfo string
	var sourceSocket string
	var targetUserInfo string
	var targetSocket string
	var parameter string
	var bigTrx string

	flag.StringVar(&sourceUserInfo, "s", "", "")
	flag.StringVar(&sourceSocket, "si", "", "")
	flag.StringVar(&targetUserInfo, "t", "", "")
	flag.StringVar(&targetSocket, "ti", "", "")
	flag.StringVar(&parameter, "c", "", "")
	flag.StringVar(&bigTrx, "m", "", "")

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
