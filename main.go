package main

import (
	"flag"
	"giogii/src/check"
	"giogii/src/flashback"
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
	var ssh string
	var sshUser string
	var sshPass string

	flag.StringVar(&sourceUserInfo, "s", "", "")
	flag.StringVar(&sourceSocket, "si", "", "")
	flag.StringVar(&targetUserInfo, "t", "", "")
	flag.StringVar(&targetSocket, "ti", "", "")
	flag.StringVar(&parameter, "c", "", "")
	flag.StringVar(&bigTrx, "m", "", "")
	flag.StringVar(&ssh, "f", "", "")
	flag.StringVar(&sshUser, "u", "", "")
	flag.StringVar(&sshPass, "p", "", "")

	flag.Parse()

	if strings.Trim(parameter, " ") == "c" {
		check.InitCheckParameterConf(sourceUserInfo, sourceSocket, "greatrds", targetUserInfo, targetSocket, "information_schema")
		check.DoCheckParameter(parameter)
	} else if strings.Trim(bigTrx, " ") == "m" {
		lock.InitConf(sourceUserInfo, sourceSocket, "performance_schema")
		lock.DoMonitorLock()
	} else if strings.Trim(ssh, " ") == "start" {
		flashback.InitMasterConnection(sourceUserInfo, sourceSocket)
		flashback.InitSlaveConnection(targetUserInfo, targetSocket)
		flashback.DoStartFlashback(targetUserInfo, targetSocket, sshUser, sshPass)
	} else if strings.Trim(ssh, " ") == "stop" {
		flashback.InitMasterConnection(sourceUserInfo, sourceSocket)
		flashback.InitSlaveConnection(targetUserInfo, targetSocket)
		flashback.DoEndFlashback(sourceUserInfo, targetUserInfo, targetSocket, sshUser, sshPass)
	} else {
		check.InitCheckConsistentConf(sourceUserInfo, sourceSocket, "information_schema", targetUserInfo, targetSocket, "information_schema")
		check.DoCheck()
	}

}
