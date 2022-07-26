package models

type GenderData string

var Gender = struct {
	Male   GenderData
	Female GenderData
}{
	Male:   "Male",
	Female: "Female",
}
