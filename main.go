package main

import (
	"flag"
	"giogii/src/check"
)

func main() {
	var sourceUserInfo string
	var sourceSocket string
	var targetUserInfo string
	var targetSocket string

	flag.StringVar(&sourceUserInfo, "s", "admin:!QAZ2wsx", "")
	flag.StringVar(&sourceSocket, "si", "172.16.128.13:16310", "")
	flag.StringVar(&targetUserInfo, "t", "admin:!QAZ2wsx", "")
	flag.StringVar(&targetSocket, "ti", "172.16.76.106:16310", "")
	flag.Parse()

	check.ConfigInit(sourceUserInfo, sourceSocket, targetUserInfo, targetSocket)
	check.DoCheck()

}
