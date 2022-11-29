package entity

import "time"

type SysInnodbLockWaits struct {
	waitStarted               time.Time
	waitAge                   time.Time
	waitAgeSecs               *int64
	lockedTable               string
	lockedTableSchema         string
	lockedTableName           string
	lockedTablePartition      string
	lockedTableSubpartition   string
	lockedIndex               string
	lockedType                string
	waitingTrxId              *int64
	waitingTrxStarted         time.Time
	waitingTrxAge             time.Time
	waitingTrxRowsLocked      *int64
	waitingTrxRowsModified    *int64
	waitingPid                *int64
	waitingQuery              string
	waitingLockId             string
	waitingLockMode           string
	blockingTrxId             *int64
	blockingPid               *int64
	blockingQuery             string
	blockingLockId            string
	blockingLockMode          string
	blockingTrxStarted        time.Time
	blockingTrxAge            time.Time
	blockingTrxRowsLocked     *int64
	blockingTrxRowsModified   *int64
	sqlKillBlockingQuery      string
	sqlKillBlockingConnection string
}
