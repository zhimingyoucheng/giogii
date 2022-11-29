package entity

type MetadataLocks struct {
	ObjectType          string
	ObjectSchema        string
	ObjectName          string
	ColumnName          string
	ObjectInstanceBegin *int64
	LockType            string
	LockDuration        string
	LockStatus          string
	Source              string
	OwnerThreadId       *int64
	OwnerEventId        *int64
}
