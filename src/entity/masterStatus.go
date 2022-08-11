package entity

type MasterStatus struct {
	file            string
	position        string
	binlogDoDB      string
	binlogIgnoreDB  string
	executedGtidSet string
}
