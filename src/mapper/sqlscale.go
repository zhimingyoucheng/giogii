package mapper

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

type SqlScaleOperator interface {
	// InitDbConnection get connection
	InitDbConnection()

	// DoQuery query
	doQuery(sqlStr string) *sql.Rows

	// DoQueryParseString return string
	DoQueryParseString(sqlStr string) string

	// DoQueryParseMap return map
	DoQueryParseMap(sqlStr string) map[string]string

	DBQueryDateString(sqlStr string) string
}

type SqlScaleStruct struct {
	DirverName       string
	ConnInfo         string
	DBconnIdleTime   time.Duration
	MaxIdleConns     int
	dbConnSocketinfo *sql.DB
}

func (sqlScaleStruct *SqlScaleStruct) DoQueryParseString(sqlStr string) string {
	//TODO implement me
	return ""
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

	stmt, err := dbConnection.Prepare(sqlStr)
	if err != nil {
		log.Println(fmt.Sprintf("Prepare SQL file ,This is a bad connection. SQL info: %s", sqlStr))
	}
	log.Println(fmt.Sprintf("Execute SQL statement queries: %s", sqlStr))

	rows, err := stmt.Query(sqlStr)
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

func (sqlScaleStruct *SqlScaleStruct) DBQueryDateString(sqlStr string) string {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var c string
	for rows.Next() {
		rows.Scan(&c)
		set := rows.NextResultSet()
		log.Println(set)
		columns, _ := rows.Columns()
		for i, v := range columns {
			log.Println(i, v)
		}

	}
	return c
}
