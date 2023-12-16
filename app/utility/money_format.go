package utility

import "github.com/leekchan/accounting"

func RupiahFormat(value int64) string {

	ac := accounting.Accounting{Symbol: "IDR ", Precision: 0}

	rupiahFormat := ac.FormatMoney(value)

	return rupiahFormat
}
