package check

import (
	"fmt"
	"giogii/src/mapper"
	"log"
	"strings"
	"time"
)

var strSql string
var MasterSqlScaleOperator mapper.SqlScaleOperator
var SlaveSqlScaleOperator mapper.SqlScaleOperator

func ConfigInit(sourceUserInfo string, sourceSocket string, targetUserInfo string, targetSocket string) {

	var masterSqlStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   1,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 3,
		ConnInfo:       fmt.Sprintf("%s@tcp(%s)/information_schema", sourceUserInfo, sourceSocket),
	}
	MasterSqlScaleOperator = masterSqlStruct

	var slaveSqlStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   1,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 3,
		ConnInfo:       fmt.Sprintf("%s@tcp(%s)/information_schema", targetUserInfo, targetSocket),
	}
	SlaveSqlScaleOperator = slaveSqlStruct
}

func DoCheck() {

	/**
	是否需要判断是否是主集群？
	是否需要判断是发是备集群？
	*/
	var rs = 2

	MasterSqlScaleOperator.InitDbConnection()
	SlaveSqlScaleOperator.InitDbConnection()

	strSql = fmt.Sprint("show master status")
	masterStatus := MasterSqlScaleOperator.DoQueryParseMaster(strSql)

	strSql = fmt.Sprint("show variables like 'server_uuid'")
	masterUuid := MasterSqlScaleOperator.DoQueryParseString(strSql)

	strSql = fmt.Sprint("show slave status")
	slaveStatus := SlaveSqlScaleOperator.DoQueryParseSlave(strSql)

	var masterGtid string
	var slaveGtid string

	if masterStatus.File != "" && slaveStatus.MasterLogFile != "" {
		masterExecutedGtids := strings.Split(masterStatus.ExecutedGtidSet, ",")
		for i := 0; i < len(masterExecutedGtids); i++ {
			if strings.Contains(masterExecutedGtids[i], masterUuid) {
				masterGtid = strings.Trim(masterExecutedGtids[i], "\n")
				break
			}
		}

		slaveExecutedGtids := strings.Split(slaveStatus.ExecutedGtidSet, ",")
		for i := 0; i < len(slaveExecutedGtids); i++ {
			if strings.Contains(slaveExecutedGtids[i], masterUuid) {
				slaveGtid = strings.Trim(slaveExecutedGtids[i], "\n")
				break
			}
		}
		if strings.Contains(slaveGtid, masterGtid) {
			rs -= 1
		}
		if masterStatus.Position == slaveStatus.ReadMasterLogPos {
			rs -= 1
		}
		log.Printf("主集群GTID：%s", masterGtid)
		log.Printf("备集群GTID：%s", slaveGtid)
		log.Print("主集群POS点位: ", masterStatus.Position)
		log.Print("备集群POS点位: ", slaveStatus.ReadMasterLogPos)

		fmt.Println(rs)

		if rs == 0 {
			MasterSqlScaleOperator.DoClose()
			SlaveSqlScaleOperator.DoClose()
		}
	}

}
