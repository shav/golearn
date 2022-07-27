package main

import (
	"fmt"
	"github.com/klauspost/lctime"
	"time"

	"github.com/bmuller/arrow"
	"github.com/goodsign/monday"
)

func main() {
	// Now
	now := time.Now()
	fmt.Println("Now: ", now)

	timestamp := time.Now().Unix()
	fmt.Println("Timestamp: ", timestamp)

	// часовые пояса
	localDateTime := time.Date(2022, time.March, 5, 8, 5, 2, 0, time.Local)
	fmt.Println(localDateTime)

	utcDateTime := time.Date(2022, time.March, 5, 8, 5, 2, 0, time.UTC)
	fmt.Println(utcDateTime)

	mskTimeZone := time.FixedZone("MSK", +3*60*60 /* sec */)
	mskDateTime := time.Date(2022, time.March, 5, 8, 5, 2, 0, mskTimeZone)
	fmt.Println(mskDateTime)

	fmt.Println("--------------------------------------")

	// форматирование дат
	fmt.Println("Standard format:")
	fmt.Println(localDateTime.Format("02 January 2006 15:04:05"))
	fmt.Println(localDateTime.Format("02.01.2006"))
	fmt.Println()

	fmt.Println("goodsign/monday format:")
	var locale monday.Locale = monday.LocaleRuRU
	fmt.Println(monday.Format(localDateTime, "02 January 2006 15:04:05", locale))
	fmt.Println(monday.Format(localDateTime, "02 January 2006 (Mon)", locale))
	fmt.Println(monday.Format(localDateTime, monday.MediumFormatsByLocale[locale], locale))
	fmt.Println()

	fmt.Println("bmuller/arrow format:")
	fmt.Println(arrow.New(localDateTime).CFormat("%d.%m.%Y %H:%M:%S"))
	fmt.Println(arrow.New(localDateTime).CFormat("%d %B %Y (%A)"))
	fmt.Println()

	fmt.Println("klauspost/lctime format:")
	loc := "ru_RU"
	fmt.Println(lctime.StrftimeLoc(loc, "%c", localDateTime))
	fmt.Println(lctime.StrftimeLoc(loc, "%d %b %Y (%a)", localDateTime))

	fmt.Println("--------------------------------------")

	// парсинг дат
	fmt.Println("Parse (standard):")
	// Если timezone не определена, тогда функция Parse возвращает время в UTC временной зоне
	t, err := time.Parse("02.01.2006", "31.07.2022")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// Если дана timezone, тогда она парсируется в данной временной зоне
	t, err = time.Parse("02.01.2006  3:04 -07:00 MST", "31.07.2022  1:25 +03:00 MSK")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// Обратите внимание, что ParseInLocation принимает последний параметр который является временной зоной.
	// В данном случае, используется локальная временная зона компьютера
	t, err = time.ParseInLocation("02.01.2006  3:04", "31.07.2022  1:25", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	fmt.Println("Parse (goodsign/monday):")
	t, err = monday.Parse("02 January 2006 15:04:05", "27 июля 2022 17:25:01", monday.LocaleRuRU)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	fmt.Println("Parse (bmuller/arrow):")
	t2, err := arrow.CParse("%d.%m.%Y %H:%M:%S", "27.07.2022 17:30:11")
	if err != nil {
		panic(err)
	}
	fmt.Println(t2)

	fmt.Println("--------------------------------------")

	// Составные части дат-времени
	t = localDateTime
	year, month, dayOfYear, dayOfMonth, weekDay := t.Year(), t.Month(), t.YearDay(), t.Day(), t.Weekday()
	hour, minute, second := t.Hour(), t.Minute(), t.Second()
	fmt.Printf("%d day of %s %d is %s\n", dayOfMonth, month, year, weekDay)
	fmt.Printf("%d day of %d year is %s\n", dayOfYear, year, t.Format("02.01.2006"))
	fmt.Printf("Hours: %d Minutes: %d Seconds: %d\n", hour, minute, second)
	fmt.Println(t.Date())
	fmt.Println(t.Clock())

	fmt.Println("--------------------------------------")

	// математика даты-времени
	origin := mskDateTime
	dateTimeNext3 := origin.Add(72 * time.Hour)
	fmt.Printf("+3 days: %v\n", dateTimeNext3)

	dateTimePrev3 := origin.Add(-72 * time.Hour)
	fmt.Printf("-3 days: %v\n", dateTimePrev3)

	fmt.Printf("-3days is before +3days: %v\n", dateTimePrev3.Before(dateTimeNext3))
	fmt.Printf("+3days is after -3days: %v\n", dateTimeNext3.After(dateTimePrev3))

	fmt.Printf("diff between +3days and origin is: %v\n", dateTimeNext3.Sub(origin))
	fmt.Printf("diff between -3days and origin is: %v\n", dateTimePrev3.Sub(origin))

	mskDateTime = time.Date(2022, time.March, 5, 8, 5, 2, 0, mskTimeZone)
	mskDateTime2 := time.Date(2022, time.March, 5, 8, 5, 2, 0, mskTimeZone)
	fmt.Println("MSK datetime equals: ", mskDateTime2.Equal(mskDateTime))

	izhTimeZone := time.FixedZone("IZH", +4*60*60 /* sec */)
	izhDateTime := time.Date(2022, time.March, 5, 9, 5, 2, 0, izhTimeZone)
	fmt.Println("MSK and IZH datetime equals: ", izhDateTime.Equal(mskDateTime))
}
