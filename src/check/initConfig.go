package check

import (
	"fmt"
	"giogii/src/mapper"
	"time"
)

func InitConfig(sourceUserInfo string, sourceSocket string, targetUserInfo string, targetSocket string) (s mapper.SqlScaleOperator, t mapper.SqlScaleOperator) {

	var masterSqlStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   1,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 1,
		ConnInfo:       fmt.Sprintf("%s@tcp(%s)/information_schema", sourceUserInfo, sourceSocket),
	}
	s = masterSqlStruct

	var slaveSqlStruct = &mapper.SqlScaleStruct{
		MaxIdleConns:   1,
		DirverName:     "mysql",
		DBconnIdleTime: time.Minute * 1,
		ConnInfo:       fmt.Sprintf("%s@tcp(%s)/information_schema", targetUserInfo, targetSocket),
	}
	t = slaveSqlStruct
	s.InitDbConnection()
	t.InitDbConnection()
	return
}
