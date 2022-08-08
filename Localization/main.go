package main

import (
	"fmt"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
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
