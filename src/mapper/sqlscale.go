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
	DoQueryParseMaster(sqlStr string) entity.MasterStatus
	DoQueryParseSlave(sqlStr string) entity.SlaveStatus
	DoQueryParseString(sqlStr string) string
}

type SqlScaleStruct struct {
	DirverName       string
	ConnInfo         string
	DBconnIdleTime   time.Duration
	MaxIdleConns     int
	dbConnSocketinfo *sql.DB
}

func (sqlScaleStruct *SqlScaleStruct) InitDbConnection() {
	db, err := sql.Open(sqlScaleStruct.DirverName, sqlScaleStruct.ConnInfo)
	if err != nil {
		errStr := fmt.Sprintf("unknown driver %q (forgotten import?)", sqlScaleStruct.DirverName)
		log.Println(errStr)
		os.Exit(1)
	}
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
	rows, err := dbConnection.Query(sqlStr)
	if err != nil {
		log.Println(fmt.Sprintf("Execute SQL file ,This is a bad connection. SQL info: %s", sqlStr))
	}
	return rows
}

func (sqlScaleStruct *SqlScaleStruct) DoQueryParseMaster(sqlStr string) entity.MasterStatus {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var masterStatus entity.MasterStatus
	for rows.Next() {
		rows.Scan(&masterStatus.File, &masterStatus.Position, &masterStatus.BinlogDoDB, &masterStatus.BinlogIgnoreDB, &masterStatus.ExecutedGtidSet)
	}
	return masterStatus
}

func (sqlScaleStruct *SqlScaleStruct) DoQueryParseSlave(sqlStr string) entity.SlaveStatus {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var slaveStatus entity.SlaveStatus
	for rows.Next() {
		err := rows.Scan(&slaveStatus.SlaveIOState, &slaveStatus.MasterHost, &slaveStatus.MasterUser, &slaveStatus.MasterPort, &slaveStatus.ConnectRetry,
			&slaveStatus.MasterLogFile, &slaveStatus.ReadMasterLogPos, &slaveStatus.RelayLogFile, &slaveStatus.RelayLogPos, &slaveStatus.RelayMasterLogFile,
			&slaveStatus.SlaveIORunning, &slaveStatus.SlaveSQLRunning, &slaveStatus.ReplicateDoDB, &slaveStatus.ReplicateIgnoreDB, &slaveStatus.ReplicateDoTable,
			&slaveStatus.ReplicateIgnoreTable, &slaveStatus.ReplicateWildDoTable, &slaveStatus.ReplicateWildIgnoreTable, &slaveStatus.LastErrno,
			&slaveStatus.LastError, &slaveStatus.SkipCounter, &slaveStatus.ExecMasterLogPos, &slaveStatus.RelayLogSpace, &slaveStatus.UntilCondition,
			&slaveStatus.UntilLogFile, &slaveStatus.UntilLogPos, &slaveStatus.MasterSSLAllowed, &slaveStatus.MasterSSLCAFile, &slaveStatus.MasterSSLCAPath,
			&slaveStatus.MasterSSLCert, &slaveStatus.MasterSSLCipher, &slaveStatus.MasterSSLKey, &slaveStatus.SecondsBehindMaster, &slaveStatus.MasterSSLVerifyServerCert,
			&slaveStatus.LastIOErrno, &slaveStatus.LastIOError, &slaveStatus.LastSQLErrno, &slaveStatus.LastSQLError, &slaveStatus.ReplicateIgnoreServerIds, &slaveStatus.MasterServerId,
			&slaveStatus.MasterUUID, &slaveStatus.MasterInfoFile, &slaveStatus.SQLDelay, &slaveStatus.SQLRemainingDelay, &slaveStatus.SlaveSQLRunningState,
			&slaveStatus.MasterRetryCount, &slaveStatus.MasterBind, &slaveStatus.LastIOErrorTimestamp, &slaveStatus.LastSQLErrorTimestamp, &slaveStatus.MasterSSLCrl,
			&slaveStatus.MasterSSLCrlpath, &slaveStatus.RetrievedGtidSet, &slaveStatus.ExecutedGtidSet, &slaveStatus.AutoPosition, &slaveStatus.ReplicateRewriteDB,
			&slaveStatus.ChannelName, &slaveStatus.MasterTLSVersion, &slaveStatus.Masterpublickeypath, &slaveStatus.Getmasterpublickey, &slaveStatus.NetworkNamespace,
		)
		if err != nil {
			log.Println(err)
		}
	}
	return slaveStatus
}

func (sqlScaleStruct *SqlScaleStruct) DoQueryParseString(sqlStr string) string {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var rst string
	var value string
	for rows.Next() {
		err := rows.Scan(&rst, &value)
		if err != nil {
			log.Println(err)
		}
	}

	return value
}
