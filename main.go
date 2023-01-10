package main

import (
	"bufio"
	"flag"
	"fmt"
	"giogii/src/check"
	"giogii/src/entity"
	"giogii/src/flashback"
	"giogii/src/lock"
	"log"
	"os"
	"strings"
)

func main() {
	var sourceUserInfo string
	var sourceSocket string
	var targetUserInfo string
	var targetSocket string
	var compare string
	var locks string
	var fb string
	var sshUser string
	var sshPass string

	flag.StringVar(&sourceUserInfo, "s", "", "")
	flag.StringVar(&sourceSocket, "si", "", "")
	flag.StringVar(&targetUserInfo, "t", "", "")
	flag.StringVar(&targetSocket, "ti", "", "")
	flag.StringVar(&compare, "c", "", "")
	flag.StringVar(&locks, "m", "", "")
	flag.StringVar(&fb, "f", "", "")
	flag.StringVar(&sshUser, "u", "", "")
	flag.StringVar(&sshPass, "p", "", "")

	flag.Parse()

	if strings.Trim(compare, " ") == "c" {
		if sourceUserInfo == "" || sourceSocket == "" || targetUserInfo == "" || targetSocket == "" {
			log.Println("parameter error")
			os.Exit(-1)
		}
		check.InitCheckParameterConf(sourceUserInfo, sourceSocket, "greatrds", targetUserInfo, targetSocket, "information_schema")
		check.DoCheckParameter(compare)
	} else if strings.Trim(locks, " ") == "m" {
		if sourceUserInfo == "" || sourceSocket == "" {
			log.Println("parameter error")
			os.Exit(-1)
		}
		lock.InitConf(sourceUserInfo, sourceSocket, "performance_schema")
		lock.DoMonitorLock()
	} else if strings.Trim(fb, " ") != "" {
		expr := strings.Trim(fb, " ")
		if sourceSocket == "" || targetSocket == "" {
			log.Println("parameter error")
			os.Exit(-1)
		}
		switch expr {
		case "verify":
			f := new(entity.FlashbackInfo)
			f.SetSourceSocket(sourceSocket)
			f.SetTargetSocket(targetSocket)
			flashback.VerifyFlashbackEnv(*f)
			// verify cluster consistent
			_, err := flashback.VerifyClusterConsistent(*f)
			if err == nil {
				log.Println("cluster consistent check success!")
			} else {
				os.Exit(-1)
			}
		case "begin":
			f := new(entity.FlashbackInfo)
			f.SetSourceSocket(sourceSocket)
			f.SetTargetSocket(targetSocket)
			flashback.VerifyFlashbackEnv(*f)
			// verify cluster consistent
			_, err := flashback.VerifyClusterConsistent(*f)
			if err == nil {
				log.Println("cluster consistent check success!")
			} else {
				os.Exit(-1)
			}
			flashback.DoBeginFlashback(f.SourceUserInfo(), f.SourceSocket(), f.TargetUserInfo(), f.TargetSocket())
		case "end":
			f := new(entity.FlashbackInfo)
			f.SetSourceSocket(sourceSocket)
			f.SetTargetSocket(targetSocket)
			flashback.VerifyFlashbackEnv(*f)
			flashback.DoEndFlashback(f.SourceUserInfo(), f.SourceSocket(), f.TargetUserInfo(), f.TargetSocket(), f.SshUser(), f.SshPass())
		case "start":
			f := new(entity.FlashbackInfo)
			f.SetSourceSocket(sourceSocket)
			f.SetTargetSocket(targetSocket)
			flashback.VerifyFlashbackEnv(*f)
			// verify cluster consistent
			_, err := flashback.VerifyClusterConsistent(*f)
			if err == nil {
				log.Println("cluster consistent check success!")
			} else {
				os.Exit(-1)
			}
			flashback.DoStartFlashback(f.SourceUserInfo(), f.SourceSocket(), f.TargetUserInfo(), f.TargetSocket(), f.SshUser(), f.SshPass())
		case "stop":
			f := new(entity.FlashbackInfo)
			f.SetSourceSocket(sourceSocket)
			f.SetTargetSocket(targetSocket)
			flashback.VerifyFlashbackEnv(*f)
			flashback.DoStopFlashback(f.SourceUserInfo(), f.SourceSocket(), f.TargetUserInfo(), f.TargetSocket(), f.SshUser(), f.SshPass())
		}
	} else {
		if sourceUserInfo == "" || sourceSocket == "" || targetUserInfo == "" || targetSocket == "" {
			log.Println("parameter error")
			os.Exit(-1)
		}
		check.InitCheckConsistentConf(sourceUserInfo, sourceSocket, "information_schema", targetUserInfo, targetSocket, "information_schema")
		check.DoCheck()
	}

}

func CallInteractive() (user string, pass string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")
		//cmd := exec.Command(cmdString)
		//cmd.Stderr = os.Stderr
		//cmd.Stdout = os.Stdout
		//err = cmd.Run()
		//if err != nil {
		//	fmt.Fprintln(os.Stderr, err)
		//}
		if user == "" {
			user = cmdString
		} else {
			pass = cmdString
			return
		}

	}
}
