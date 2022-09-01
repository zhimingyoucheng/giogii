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

	flag.StringVar(&sourceUserInfo, "s", "admin:!QAZ2wsx", "")
	flag.StringVar(&sourceSocket, "si", "172.17.128.151:16310", "")
	flag.StringVar(&targetUserInfo, "t", "admin:!QAZ2wsx", "")
	flag.StringVar(&targetSocket, "ti", "172.17.128.13:16310", "")
	flag.StringVar(&parameter, "c", "parameters", "")
	flag.Parse()

	check.ConfigInit(sourceUserInfo, sourceSocket, targetUserInfo, targetSocket)

	if strings.Trim(parameter, " ") == "parameter" {
		check.DoCheckParameter()
	} else {
		check.DoCheck()
	}

}
