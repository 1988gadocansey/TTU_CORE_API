package models

type PaymentSources string

var PaymentSource = struct {
	Bank        PaymentSources
	Momo        PaymentSources
	ExcelUpload PaymentSources
}{
	Bank:        "Bank",
	Momo:        "Momo",
	ExcelUpload: "ExcelUpload",
}
