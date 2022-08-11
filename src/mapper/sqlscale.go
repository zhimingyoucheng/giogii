package mapper

import (
	"database/sql"
	"fmt"
	"giogii/src/entity"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

type SqlScaleOperator interface {
	InitDbConnection()
	doQuery(sqlStr string) *sql.Rows
	DoQueryParseString(sqlStr string) entity.MasterStatus
	DoQueryParseMap(sqlStr string) map[string]string
}

type SqlScaleStruct struct {
	DirverName       string
	ConnInfo         string
	DBconnIdleTime   time.Duration
	MaxIdleConns     int
	dbConnSocketinfo *sql.DB
}

func (sqlScaleStruct *SqlScaleStruct) InitDbConnection() {
	log.Println("Initializes the database connection object")
	db, err := sql.Open(sqlScaleStruct.DirverName, sqlScaleStruct.ConnInfo)
	if err != nil {
		errStr := fmt.Sprintf("unknown driver %q (forgotten import?)", sqlScaleStruct.DirverName)
		log.Println(errStr)
		os.Exit(1)
	}
	log.Println("Send a ping packet to check the database running status")
	if err := db.Ping(); err != nil {
		errStr := "Failed to open a database connection and create a session connection. pleace Check the database status or network status"
		log.Println(errStr)
		os.Exit(1)
	}
	db.SetConnMaxIdleTime(sqlScaleStruct.DBconnIdleTime)
	db.SetMaxIdleConns(sqlScaleStruct.MaxIdleConns)
	sqlScaleStruct.dbConnSocketinfo = db
}

func (sqlScaleStruct *SqlScaleStruct) doQuery(sqlStr string) *sql.Rows {
	dbConnection := sqlScaleStruct.dbConnSocketinfo
	log.Println(fmt.Sprintf("Prepare initialize the SQL statement:%s", sqlStr))

	/*	stmt, err := dbConnection.Prepare(sqlStr)
		if err != nil {
			log.Println(fmt.Sprintf("Prepare SQL file ,This is a bad connection. SQL info: %s", sqlStr))
		}
		log.Println(fmt.Sprintf("Execute SQL statement queries: %s", sqlStr))
	*/
	rows, err := dbConnection.Query(sqlStr)
	if err != nil {
		log.Println(fmt.Sprintf("Execute SQL file ,This is a bad connection. SQL info: %s", sqlStr))
	}
	return rows
}

func (sqlScaleStruct *SqlScaleStruct) DoQueryParseMap(sqlStr string) map[string]string {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var result map[string]string
	result = make(map[string]string)
	var keySlice, valSlice []string
	var v, c string
	for rows.Next() {
		rows.Scan(&v, &c)
		keySlice = append(keySlice, v)
		valSlice = append(valSlice, c)
	}
	for i := range keySlice {
		result[keySlice[i]] = valSlice[i]
	}
	return result
}

func (sqlScaleStruct *SqlScaleStruct) DoQueryParseString(sqlStr string) entity.MasterStatus {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var masterStatus entity.MasterStatus
	for rows.Next() {
		rows.Scan(&masterStatus.File, &masterStatus.Position, &masterStatus.BinlogDoDB, &masterStatus.BinlogIgnoreDB, &masterStatus.ExecutedGtidSet)
	}
	return masterStatus
}
