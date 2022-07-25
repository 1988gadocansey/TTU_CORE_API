package models

type AppTitle string

var Title = struct {
	Mrs  AppTitle
	Miss AppTitle
	Prof AppTitle
	Dr   AppTitle
	Mr   AppTitle
	Eng  AppTitle
	Rev  AppTitle
}{
	Mrs:  "Mrs",
	Miss: "Miss",
	Prof: "Prof",
	Dr:   "Dr",
	Mr:   "Mr",
	Rev:  "Rev",
	Eng:  "Eng",
}
