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

	flag.StringVar(&sourceUserInfo, "s", "", "")
	flag.StringVar(&sourceSocket, "si", "", "")
	flag.StringVar(&targetUserInfo, "t", "", "")
	flag.StringVar(&targetSocket, "ti", "", "")
	flag.StringVar(&parameter, "c", "", "")

	flag.Parse()

	if strings.Trim(parameter, " ") != "" {
		check.InitCheckParameterConf(sourceUserInfo, sourceSocket, "greatrds", targetUserInfo, targetSocket, "information_schema")
		check.DoCheckParameter(parameter)
	} else {
		check.InitCheckConsistentConf(sourceUserInfo, sourceSocket, "information_schema", targetUserInfo, targetSocket, "information_schema")
		check.DoCheck()
	}

}
