package models

type Statuses string

var Status = struct {
	Dead       Statuses
	Student    Statuses
	Abandon    Statuses
	Alumni     Statuses
	Deferred   Statuses
	Rusticated Statuses
}{
	Dead:       "Dead",
	Student:    "Student",
	Abandon:    "Abandon",
	Deferred:   "Deferred",
	Rusticated: "Rusticated",
}
