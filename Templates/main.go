package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Age     int
	Manager *Person
}

func main() {
	// Парсинг шаблона
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		fmt.Println(err)
	}
	// Валидация шаблона (генерирует panic при наличии ошибок в шаблоне!)
	template.Must(t1, err)

	fmt.Println("----------------------------------")

	// Подстановка в шаблон элементарных значений
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)

	list := []string{"Go", "Rust", "C++", "C#"}
	t1.Execute(os.Stdout, list)

	fmt.Println("----------------------------------")

	t2 := CreateTemplate("t2", "Name is {{.Name}}\n", nil)

	// Подстановка в шаблон структур
	person := &Person{Name: "Artem", Age: 20}
	t2.Execute(os.Stdout, person)

	myMap := map[string]string{
		"Name": "Tom",
	}
	t2.Execute(os.Stdout, myMap)

	t20 := CreateTemplate("t20", "Manager is {{.Manager.Name}}\n", nil)
	t20.Execute(os.Stdout, person)
	fmt.Println()
	manager := &Person{Name: "Tom"}
	person.Manager = manager
	t20.Execute(os.Stdout, person)

	fmt.Println("----------------------------------")

	// Условия
	t3 := CreateTemplate("t3",
		"{{if . -}} {{.}} {{else -}} Data is empty {{end}}\n", nil)
	t3.Execute(os.Stdout, "not empty string")
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, list)
	t3.Execute(os.Stdout, []string{})

	fmt.Println("----------------------------------")

	// Перечисления
	t4 := CreateTemplate("t4",
		"Items: {{range .}}{{.}}, {{end}}\n", nil)
	t4.Execute(os.Stdout, list)

	fmt.Println("----------------------------------")

	// Кастомные функции
	t5 := CreateTemplate("t5",
		"NameAge is {{concat .Name .Age}}\n",
		template.FuncMap{"concat": concat})
	t5.Execute(os.Stdout, person)
}

func CreateTemplate(name string, tpl string, funcs template.FuncMap) *template.Template {
	t := template.New(name)
	if funcs != nil {
		t = t.Funcs(funcs)
	}
	return template.Must(t.Parse(tpl))
}

func concat(name string, age int) string {
	return fmt.Sprintf("%s+%d", name, age)
}
