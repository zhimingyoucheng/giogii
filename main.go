package main

import (
	"flag"
	"giogii/src/check"
	"strings"
)

func main() {
	var sourceUserInfo string
	var sourceSocket string
	var targetUserInfo string
	var targetSocket string
	var parameter string

	flag.StringVar(&sourceUserInfo, "s", "root:drACgwoqtM", "")
	flag.StringVar(&sourceSocket, "si", "172.17.128.49:13336", "")
	flag.StringVar(&targetUserInfo, "t", "admin:!QAZ2wsx", "")
	flag.StringVar(&targetSocket, "ti", "172.17.128.13:16310", "")
	flag.StringVar(&parameter, "c", "", "")
	flag.Parse()

	if strings.Trim(parameter, " ") == "base" {
		check.InitCheckParameterConf(sourceUserInfo, sourceSocket, "greatrds", targetUserInfo, targetSocket, "information_schema")
		check.DoCheckParameter()
	} else {
		check.InitCheckConsistentConf(sourceUserInfo, sourceSocket, "information_schema", targetUserInfo, targetSocket, "information_schema")
		check.DoCheck()
	}

}
