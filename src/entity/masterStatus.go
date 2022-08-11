package entity

type MasterStatus struct {
	File            string
	Position        string
	BinlogDoDB      string
	BinlogIgnoreDB  string
	ExecutedGtidSet string
}
