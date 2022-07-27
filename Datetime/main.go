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
	localDateTime1 := time.Date(2022, time.March, 5, 8, 5, 2, 0, time.Local)
	fmt.Println(localDateTime1)

	utcDateTime1 := time.Date(2022, time.March, 5, 8, 5, 2, 0, time.UTC)
	fmt.Println(utcDateTime1)

	mskTimeZone := time.FixedZone("MSK", +3*60*60 /* sec */)
	mskDateTime1 := time.Date(2022, time.March, 5, 8, 5, 2, 0, mskTimeZone)
	fmt.Println(mskDateTime1)

	fmt.Println("--------------------------------------")

	// форматирование дат
	fmt.Println("Standard format:")
	fmt.Println(localDateTime1.Format("02 January 2006 15:04:05"))
	fmt.Println(localDateTime1.Format("02.01.2006"))
	fmt.Println()

	fmt.Println("goodsign/monday format:")
	var locale monday.Locale = monday.LocaleRuRU
	fmt.Println(monday.Format(localDateTime1, "02 January 2006 15:04:05", locale))
	fmt.Println(monday.Format(localDateTime1, "02 January 2006 (Mon)", locale))
	fmt.Println(monday.Format(localDateTime1, monday.MediumFormatsByLocale[locale], locale))
	fmt.Println()

	fmt.Println("bmuller/arrow format:")
	fmt.Println(arrow.New(localDateTime1).CFormat("%d.%m.%Y %H:%M:%S"))
	fmt.Println(arrow.New(localDateTime1).CFormat("%d %B %Y (%A)"))
	fmt.Println()

	fmt.Println("klauspost/lctime format:")
	loc := "ru_RU"
	fmt.Println(lctime.StrftimeLoc(loc, "%c", localDateTime1))
	fmt.Println(lctime.StrftimeLoc(loc, "%d %b %Y (%a)", localDateTime1))

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
}
