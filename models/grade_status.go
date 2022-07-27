package models

type GradeStatuses string

var GradeStatus = struct {
	Cancelled GradeStatuses
	Pending   GradeStatuses
	Released  GradeStatuses
}{
	Cancelled: "Cancelled",
	Pending:   "Pending",
	Released:  "Released",
}
