package models

type MaritalStatuses string

var MaritalStatus = struct {
	Married   MaritalStatuses
	Divorced  MaritalStatuses
	Single    MaritalStatuses
	Seperated MaritalStatuses
	Widow     MaritalStatuses
	Widower   MaritalStatuses
}{
	Married:   "Married",
	Divorced:  "Divorced",
	Single:    "Single",
	Seperated: "Seperated",
	Widow:     "Widow",
	Widower:   "Widower",
}
