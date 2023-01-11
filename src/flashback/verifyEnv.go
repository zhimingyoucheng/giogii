package flashback

import (
	"errors"
	"fmt"
	"giogii/src/entity"
	"giogii/src/file"
	"giogii/src/mapper"
	"log"
	"strings"
)

func VerifyFlashbackEnv(f entity.FlashbackInfo) (*entity.FlashbackInfo, error) {

	// 1) read config file
	fileName := "gii.conf"
	f, _ = file.ReadConfig(fileName, f)
	log.Printf("%s check success!", fileName)

	// 2) verify source and target parameter
	_, err := verifyDbConn(f.SourceUserInfo(), f.SourceSocket())
	if err == nil {
		log.Println("sourceDataSource check success!")
	} else {
		log.Println("sourceDataSource check failed!")
	}

	_, err = verifyDbConn(f.TargetUserInfo(), f.TargetSocket())
	if err == nil {
		log.Println("targetDataSource check success!")
	} else {
		log.Println("targetDataSource check failed!")
	}

	_, err = verifySsh(f.TargetUserInfo(), f.TargetSocket(), f.SshUser(), f.SshPass(), f.SshPort())
	if err == nil {
		log.Println("ssh check success!")
	} else {
		log.Println(err)
		log.Println("ssh check failed!")
	}

	return &f, nil
}

// 校验B集群是否是A集群的备集群
// TODO

func VerifyReplicationConsistent(f entity.FlashbackInfo) (string, error) {
	var sqlScale mapper.SqlScaleOperator
	defer func() {
		sqlScale.DoClose()
	}()
	conn := mapper.CreateConn(f.TargetUserInfo(), f.TargetSocket(), "information_schema")
	sqlScale = &conn
	strSql := fmt.Sprint("dbscale show dataservers")
	m := sqlScale.DoQueryParseToDataServers(strSql)
	if len(m) == 0 {
		return "", errors.New("fail")
	}

	for i := 0; i < len(m); i++ {
		if strings.Contains(m[i].Servername.String, "server") { // contains backup cluster server
			status := m[i].Status.String
			if status != "" && strings.Contains(status, "down") {
				return "", errors.New("fail")
			}
		}
	}
	return "", nil
}

func VerifyClusterConsistent(f entity.FlashbackInfo) (string, error) {
	var sqlScale mapper.SqlScaleOperator
	defer func() {
		sqlScale.DoClose()
	}()
	conn := mapper.CreateConn(f.TargetUserInfo(), f.TargetSocket(), "information_schema")
	sqlScale = &conn
	ips := getIp(&sqlScale)

	var s []map[string]string

	for i := 0; i < len(ips); i++ {
		// connect 16315
		ip := ips[i]["ip"]
		port := ips[i]["port"]
		newTargetSocket := ip + ":" + port
		createConn := mapper.CreateConn(f.TargetUserInfo(), newTargetSocket, "information_schema")

		// get node binlog status
		strSql := fmt.Sprint("show master status")
		masterStatus := createConn.DoQueryParseMaster(strSql)

		createConn.DoClose()

		// get global transaction id
		globalTransactionId := strings.ReplaceAll(masterStatus.ExecutedGtidSet, "\n", "")
		myMap := make(map[string]string)
		myMap["ip"] = ip
		myMap["id"] = globalTransactionId
		s = append(s, myMap)
	}

	switch len(s) {
	case 2:
		if s[0]["id"] != s[1]["id"] {
			return "", errors.New(fmt.Sprintf("%s node compare %s node error", s[0]["ip"], s[1]["ip"]))
		}
	case 3:
		if s[0]["id"] != s[1]["id"] || s[1]["id"] != s[2]["id"] || s[0]["id"] != s[2]["id"] {
			return "", errors.New(fmt.Sprintf("%s node compare %s or %s node error", s[0]["ip"], s[1]["ip"], s[2]["ip"]))
		}
	}

	return "", nil
}

func verifyDbConn(userInfo string, socket string) (string, error) {
	var sqlScale mapper.SqlScaleOperator
	defer func() {
		sqlScale.DoClose()
	}()
	conn := mapper.InitSourceConn(userInfo, socket, "information_schema")
	sqlScale = &conn
	sqlStr := "select 1"
	value := sqlScale.DoQueryParseSingleValue(sqlStr)
	if value == "1" {
		return "", nil
	} else {
		return "", errors.New("fail")
	}
}

func getIp(sqlScale *mapper.SqlScaleOperator) (s []map[string]string) {
	strSql := fmt.Sprint("dbscale show dataservers")
	scale := *sqlScale
	m := scale.DoQueryParseToDataServers(strSql)
	if len(m) > 0 {
		for i := 0; i < len(m); i++ {
			myMap := make(map[string]string)
			myMap["ip"] = m[i].Host.String
			myMap["port"] = m[i].Port.String
			s = append(s, myMap)
		}
	}
	return s
}

func verifySsh(userInfo string, socket string, sshUser string, sshPass string, sshPort string) (string, error) {
	var sqlScale mapper.SqlScaleOperator
	defer func() {
		sqlScale.DoClose()
	}()
	conn := mapper.CreateConn(userInfo, socket, "information_schema")
	sqlScale = &conn
	ips := getIp(&sqlScale)

	for i := 0; i < len(ips); i++ {
		cli := Client{
			Username: sshUser,
			Password: sshPass,
			Socket:   fmt.Sprintf("%s:%s", ips[i]["ip"], sshPort),
		}
		_, err := cli.Connect()
		if err != nil {
			return "", errors.New(fmt.Sprintf("ssh ip %s check failed!", ips[i]))
		}
	}

	return "", nil

}
