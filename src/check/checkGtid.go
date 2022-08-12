package check

import (
	"fmt"
	"giogii/src/mapper"
	"log"
	"strings"
	"time"
)

var strSql string
var ShowMasterStatus = make(map[string]string)
var MasterSqlScaleOperator mapper.SqlScaleOperator
var SlaveSqlScaleOperator mapper.SqlScaleOperator

func ConfigInit() {
	var sqlScaleStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   2,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 3,
		ConnInfo:       "admin:!QAZ2wsx@tcp(172.16.76.105:16310)/information_schema",
	}
	MasterSqlScaleOperator = sqlScaleStruct

	var slaveSqlStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   2,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 3,
		ConnInfo:       "admin:!QAZ2wsx@tcp(172.16.128.13:16310)/information_schema",
	}
	SlaveSqlScaleOperator = slaveSqlStruct
}

func DoCheck() {

	/**
	需要判断是否是主集群
	需要判断是发是备集群
	*/
	var rs int = 2

	strSql = fmt.Sprint("show master status")
	MasterSqlScaleOperator.InitDbConnection()
	masterStatus := MasterSqlScaleOperator.DoQueryParseMaster(strSql)

	strSql = fmt.Sprint("select uuid()")
	masterUuid := MasterSqlScaleOperator.DoQueryParseString(strSql)

	strSql = fmt.Sprint("show slave status")
	SlaveSqlScaleOperator.InitDbConnection()
	slaveStatus := SlaveSqlScaleOperator.DoQueryParseSlave(strSql)

	strSql = fmt.Sprint("select uuid()")
	//slaveUuid := SlaveSqlScaleOperator.DoQueryParseString(strSql)

	var masterGtid string
	var slaveGtid string

	if masterStatus.File != "" && slaveStatus.MasterLogFile != "" {
		masterExectedGtids := strings.Split(masterStatus.ExecutedGtidSet, ",")
		for i := 0; i < len(masterExectedGtids); i++ {
			log.Println(masterExectedGtids[i])
			if strings.Contains(masterExectedGtids[i], masterUuid) {
				masterGtid = masterExectedGtids[i]
				break
			}
		}

		slaveExectedGtids := strings.Split(slaveStatus.ExecutedGtidSet, ",")
		for i := 0; i < len(slaveExectedGtids); i++ {
			log.Println(slaveExectedGtids[i])
			if strings.Contains(slaveExectedGtids[i], masterUuid) {
				slaveGtid = slaveExectedGtids[i]
				break
			}
		}

		if strings.Contains(slaveGtid, masterGtid) {
			rs -= 1
		}

	}

}
