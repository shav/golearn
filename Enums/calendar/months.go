package calendar

import (
	"errors"
	"strings"
)

type Month string

type Year struct {
	Unknown   Month
	January   Month
	February  Month
	March     Month
	April     Month
	May       Month
	June      Month
	July      Month
	August    Month
	September Month
	October   Month
	November  Month
	December  Month
}

var Months = Year{
	January:   "january",
	February:  "february",
	March:     "march",
	April:     "april",
	May:       "may",
	June:      "june",
	July:      "july",
	August:    "august",
	September: "september",
	October:   "october",
	November:  "november",
	December:  "december",
}

var months = map[string]Month{
	string(Months.January):   Months.January,
	string(Months.February):  Months.February,
	string(Months.March):     Months.March,
	string(Months.April):     Months.April,
	string(Months.May):       Months.May,
	string(Months.June):      Months.June,
	string(Months.July):      Months.July,
	string(Months.August):    Months.August,
	string(Months.September): Months.September,
	string(Months.October):   Months.October,
	string(Months.November):  Months.November,
	string(Months.December):  Months.December,
}

func (month Month) String() string {
	return string(month)
}

func MonthFromString(str string) (Month, error) {
	if month, ok := months[strings.ToLower(str)]; ok {
		return month, nil
	}
	return Months.Unknown, errors.New("unknown month")
}
