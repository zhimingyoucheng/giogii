package mapper

import (
	"fmt"
	"time"
)

func InitConfig(sourceUserInfo string, sourceSocket string, sourceDatabase string, targetUserInfo string, targetSocket string, targetDatabase string) (s SqlStruct, t SqlStruct) {

	s = SqlStruct{
		MaxIdleConn:  1,
		DriverName:   "mysql",
		ConnIdleTime: time.Minute * 1,
		ConnInfo:     fmt.Sprintf("%s@tcp(%s)/%s", sourceUserInfo, sourceSocket, sourceDatabase),
	}

	t = SqlStruct{
		MaxIdleConn:  1,
		DriverName:   "mysql",
		ConnIdleTime: time.Minute * 1,
		ConnInfo:     fmt.Sprintf("%s@tcp(%s)/%s", targetUserInfo, targetSocket, targetDatabase),
	}

	s.InitConnection()
	t.InitConnection()
	return
}
