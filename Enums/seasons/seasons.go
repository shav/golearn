package seasons

import (
	"errors"
	"strings"
)

type Season string

const (
	Unknown        = ""
	Winter  Season = "winter"
	Spring         = "spring"
	Summer         = "summer"
	Autumn         = "autumn"
)

func (season Season) String() string {
	return string(season)
}

func FromString(str string) (Season, error) {
	switch strings.ToLower(str) {
	case string(Winter):
		return Winter, nil
	case string(Spring):
		return Spring, nil
	case string(Summer):
		return Summer, nil
	case string(Autumn):
		return Autumn, nil
	default:
		return Unknown, errors.New("unknown season")
	}
}
