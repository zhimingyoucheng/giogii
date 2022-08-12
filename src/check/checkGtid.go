package check

import (
	"fmt"
	"giogii/src/mapper"
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

	strSql = fmt.Sprint("show master status")
	MasterSqlScaleOperator.InitDbConnection()
	masterStatus := MasterSqlScaleOperator.DoQueryParseMaster(strSql)
	if masterStatus.File != "" {

	}

	strSql = fmt.Sprint("show slave status")
	SlaveSqlScaleOperator.InitDbConnection()
	slaveStatus := SlaveSqlScaleOperator.DoQueryParseSlave(strSql)

	if slaveStatus.MasterLogFile != "" {

	}
}
