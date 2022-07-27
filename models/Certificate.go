package models

type Awards string

var Award = struct {
	Bachelor    Awards
	Diploma     Awards
	Masters     Awards
	Doctorate   Awards
	Certificate Awards
}{
	Bachelor:    "Bachelor",
	Diploma:     "Diploma",
	Doctorate:   "Doctorate",
	Masters:     "Masters",
	Certificate: "Certificate",
}
