package mapper

import (
	"fmt"
	"time"
)

func InitConfig(sourceUserInfo string, sourceSocket string, targetUserInfo string, targetSocket string) (s SqlStruct, t SqlStruct) {

	s = SqlStruct{
		MaxIdleConn:  1,
		DriverName:   "mysql",
		ConnIdleTime: time.Minute * 1,
		ConnInfo:     fmt.Sprintf("%s@tcp(%s)/information_schema", sourceUserInfo, sourceSocket),
	}

	t = SqlStruct{
		MaxIdleConn:  1,
		DriverName:   "mysql",
		ConnIdleTime: time.Minute * 1,
		ConnInfo:     fmt.Sprintf("%s@tcp(%s)/information_schema", targetUserInfo, targetSocket),
	}

	s.InitConnection()
	t.InitConnection()
	return
}
