package check

import (
	"fmt"
	"giogii/src/mapper"
	"time"
)

var strSql string
var ShowMasterStatus = make(map[string]string)
var SqlScaleOperator mapper.SqlScaleOperator

func ConfigInit() {
	var sqlScaleStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   2,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 3,
		ConnInfo:       "admin:!QAZ2wsx@tcp(172.16.76.105:16310)/information_schema",
	}
	SqlScaleOperator = sqlScaleStruct
}

func DoCheck() {

	strSql = fmt.Sprint("show master status")
	SqlScaleOperator.InitDbConnection()
	masterStatus := SqlScaleOperator.DoQueryParseString(strSql)
	if masterStatus.File != "" {

	}

	strSql = fmt.Sprint("show slave status")
	SqlScaleOperator.InitDbConnection()

}
