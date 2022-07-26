package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/messages.en.json")
	bundle.LoadMessageFile("resources/messages.ru.json")
}

func main() {
	localizer := i18n.NewLocalizer(bundle, language.Russian.String())
	lStr, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "welcome",
	})
	fmt.Println(lStr)

	lStr2, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "hello",
		TemplateData: map[string]string{
			"Name":  "Artem",
			"Today": time.Now().Format("02.01.2006"),
		},
	})
	fmt.Println(lStr2)

	lStr3, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:   "ItemsCount",
		PluralCount: 2,
	})
	fmt.Println(lStr3)

	// Множественное число (фигня полная, для русского языка в таком виде это не подходит!)
	lStr4, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: "Cookies"},
		PluralCount:    1,
	})
	fmt.Println(lStr4)

	fmt.Println("--------------------------------------")

	// Множественное число
	const ITEMS_MESSAGE string = "%d items"
	message.Set(language.English, ITEMS_MESSAGE,
		plural.Selectf(1, "%d",
			"=0", "no items",
			plural.One, "one item",
			"<100", "%[1]d items",
			plural.Other, "a lot of items",
		),
	)

	const AVERAGE_MESSAGE string = "The average is %0.2f"
	message.Set(language.English, AVERAGE_MESSAGE,
		plural.Selectf(1, "%0.2f",
			"<1", "The average is zero",
			"=1", "The average is one",
			plural.Other, "The average is %[1]f",
		),
	)

	m := message.NewPrinter(language.English)
	fmt.Println(m.Sprintf(ITEMS_MESSAGE, 0))
	fmt.Println(m.Sprintf(ITEMS_MESSAGE, 1))
	fmt.Println(m.Sprintf(ITEMS_MESSAGE, 10))
	fmt.Println(m.Sprintf(ITEMS_MESSAGE, 1000))
	fmt.Println(m.Sprintf(AVERAGE_MESSAGE, 0.8))
	fmt.Println(m.Sprintf(AVERAGE_MESSAGE, 1.0))
	fmt.Println(m.Sprintf(AVERAGE_MESSAGE, 10.23))
}
