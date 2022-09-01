package mapper

import (
	"giogii/src/entity"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type SqlScaleOperator interface {
	DoQueryParseMaster(sqlStr string) entity.MasterStatus
	DoQueryParseSlave(sqlStr string) entity.SlaveStatus
	DoQueryParseString(sqlStr string) string
	DoClose()
}

func (sqlScaleStruct *SqlStruct) DoQueryParseMaster(sqlStr string) entity.MasterStatus {
	rows := sqlScaleStruct.doQuery(sqlStr)
	var masterStatus entity.MasterStatus
	for rows.Next() {
		rows.Scan(&masterStatus.File, &masterStatus.Position, &masterStatus.BinlogDoDB, &masterStatus.BinlogIgnoreDB, &masterStatus.ExecutedGtidSet)
	}
	return masterStatus
}

func (sqlScaleStruct *SqlStruct) DoQueryParseSlave(sqlStr string) entity.SlaveStatus {
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

func (sqlScaleStruct *SqlStruct) DoQueryParseString(sqlStr string) string {
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

func (sqlScaleStruct *SqlStruct) DoClose() {
	sqlScaleStruct.Connection.Close()
}