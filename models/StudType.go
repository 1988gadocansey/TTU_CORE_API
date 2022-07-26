package models

type StudTypes string

var StudType = struct {
	Undergraduate StudTypes
	Certificate   StudTypes
	Diploma       StudTypes
	Postgraduate  StudTypes
}{
	Undergraduate: "Undergraduate",
	Certificate:   "Certificate",
	Diploma:       "Diploma",
	Postgraduate:  "Postgraduate",
}
