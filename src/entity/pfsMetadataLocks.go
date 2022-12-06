package entity

type MetadataLocks struct {
	LockType         string
	LockStatus       string
	ProcesslistId    *int64
	ProcesslistState string
	ProcesslistInfo  string
}
