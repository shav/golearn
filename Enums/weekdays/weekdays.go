package weekdays

import (
	"errors"
	"strings"
)

type Weekday int

const (
	Unknown         = -1
	Monday  Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var weekdayNames = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (weekday Weekday) String() string {
	return weekdayNames[weekday-1]
}

//func (weekday Weekday) String() string {
//	switch weekday {
//	case Monday:
//		return "Monday"
//	case Tuesday:
//		return "Tuesday"
//	case Wednesday:
//		return "Wednesday"
//	case Thursday:
//		return "Thursday"
//	case Friday:
//		return "Friday"
//	case Saturday:
//		return "Saturday"
//	case Sunday:
//		return "Sunday"
//	default:
//		return ""
//	}
//}

func FromString(str string) (Weekday, error) {
	switch strings.ToLower(str) {
	case "monday":
		return Monday, nil
	case "tuesday":
		return Tuesday, nil
	case "wednesday":
		return Wednesday, nil
	case "thursday":
		return Thursday, nil
	case "friday":
		return Friday, nil
	case "saturday":
		return Saturday, nil
	case "sunday":
		return Sunday, nil
	default:
		return Unknown, errors.New("unknown weekday")
	}
}
