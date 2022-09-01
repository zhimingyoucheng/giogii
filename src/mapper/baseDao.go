package mapper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type SqlStruct struct {
	DriverName   string
	ConnInfo     string
	ConnIdleTime time.Duration
	MaxIdleConn  int
	Connection   *sql.DB
}

func (sqlScaleStruct *SqlStruct) InitConnection() {
	db, err := sql.Open(sqlScaleStruct.DriverName, sqlScaleStruct.ConnInfo)
	if err != nil {
		log.Printf("unknown driver %s (forgotten import?)", sqlScaleStruct.DriverName)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		log.Printf("Failed to open a database connection and create a session connection %s", sqlScaleStruct.ConnInfo)
		os.Exit(1)
	}
	db.SetConnMaxIdleTime(sqlScaleStruct.ConnIdleTime)
	db.SetMaxIdleConns(sqlScaleStruct.MaxIdleConn)
	sqlScaleStruct.Connection = db
}

func (sqlScaleStruct *SqlStruct) doQuery(sqlStr string) *sql.Rows {
	con := sqlScaleStruct.Connection
	rows, err := con.Query(sqlStr)
	if err != nil {
		log.Println(fmt.Sprintf("This is a bad connection. SQL info: %s", sqlStr))
	}
	return rows
}
