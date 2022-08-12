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

	flag.StringVar(&sourceUserInfo, "s", "", "")
	flag.StringVar(&sourceSocket, "sip", "", "")
	flag.StringVar(&targetUserInfo, "t", "", "")
	flag.StringVar(&targetSocket, "tip", "", "")
	flag.Parse()
	//fmt.Println(sourceUserInfo, sourceSocket, targetUserInfo, targetSocket)
	check.ConfigInit(sourceUserInfo, sourceSocket, targetUserInfo, targetSocket)
	check.DoCheck()

}
