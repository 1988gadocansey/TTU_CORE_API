package models

type ResitStatuses string

var ResitStatus = struct {
	No   ResitStatuses
	Yes  ResitStatuses
	Done ResitStatuses
}{
	No:   "No",
	Yes:  "Yes",
	Done: "Done",
}
