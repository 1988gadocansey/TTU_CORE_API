package models

type Sections string

var Section = struct {
	Regular  Sections
	Evening  Sections
	Weekend  Sections
	Distance Sections
	Sandwich Sections
}{
	Regular:  "Regular",
	Evening:  "Evening",
	Weekend:  "Weekend",
	Distance: "Distance",
	Sandwich: "Sandwich",
}
