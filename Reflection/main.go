// https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7

package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type Temperature int

type IAlive interface {
	Eat()
}

type Person struct {
	Name string `tag1:"First Tag" tag2:"Second Tag"`
	Age  int
}

func (person Person) Eat() {
	fmt.Println(person.Name + " eats")
}

type Plane struct {
	Model string
}

type List[T any] []T

func main() {
	// Получение типов
	str := "hello"
	strType := reflect.TypeOf(str)
	fmt.Printf("Type: %s (kind: %s)\n", strType.Name(), strType.Kind())

	pStr := &str
	pStrType := reflect.TypeOf(pStr)
	fmt.Printf("Type: %s (kind: %s to %s)\n", pStrType.Name(), pStrType.Kind(), pStrType.Elem())

	var temperature Temperature = +10
	temperatureType := reflect.TypeOf(temperature)
	fmt.Printf("Type: %s (kind: %s)\n", temperatureType.Name(), temperatureType.Kind())

	var slice = []int{1, 2, 3}
	sliceType := reflect.TypeOf(slice)
	fmt.Printf("Type: %s (kind: %s of %s)\n", sliceType.Name(), sliceType.Kind(), sliceType.Elem())

	var myMap = map[string]int{"Tom": 1}
	mapType := reflect.TypeOf(myMap)
	fmt.Printf("Type: %s (kind: %s of %s)\n", mapType.Name(), mapType.Kind(), mapType.Elem())

	var list = List[string]{}
	listType := reflect.TypeOf(list)
	fmt.Printf("Type: %s (kind: %s)\n", listType.Name(), listType.Kind())

	var person = Person{Name: "Artem", Age: 20}
	personType := reflect.TypeOf(person)
	fmt.Printf("Type: %s (kind: %s)\n", personType.Name(), personType.Kind())

	// Получение полей структуры
	for i := 0; i < personType.NumField(); i++ {
		property := personType.Field(i)
		fmt.Printf("    Property: %s type: %s (kind: %s)\n", property.Name, property.Type.Name(), property.Type.Kind())
		// Атрибуты полей
		if property.Tag != "" {
			fmt.Printf("        Tag: %s\n", property.Tag)
			fmt.Printf("            tag1: %s, tag2: %s\n", property.Tag.Get("tag1"), property.Tag.Get("tag2"))
		}
	}

	fmt.Println("--------------------------------------")

	// Получение значения переменных
	strVal := reflect.ValueOf(str)
	fmt.Println(strVal.Interface())

	// Получение значения указателя
	pStrVal := reflect.ValueOf(&str)
	fmt.Println(pStrVal.Interface()) // Адрес указателя
	fmt.Println(pStrVal.Elem())      // Объект, на который указывает указатель

	// Получение значения полей структуры
	personVal := reflect.ValueOf(&person)
	strProperty := personVal.Elem().Field(0)
	fmt.Println(strProperty.Interface())

	fmt.Println("--------------------------------------")

	// Означивание переменных
	// strVal.SetString("World") // error: Просто переменные (не указатели) изменять через reflection нельзя
	pStrVal.Elem().SetString("World")
	fmt.Println(str)

	// Означивание полей структуры
	strProperty.SetString("Tom")
	fmt.Println(person)

	fmt.Println("--------------------------------------")

	// Проверка совместимости типов между собой
	personType = TypeOf[Person]()
	planeType := TypeOf[Plane]()
	aliveInterface := TypeOf[IAlive]()
	fmt.Printf("Person implements IAlive: %t\n", personType.Implements(aliveInterface))
	fmt.Printf("Plane implements IAlive: %t\n", planeType.Implements(aliveInterface))

	intType := TypeOf[int]()
	temperatureType = TypeOf[Temperature]()
	fmt.Printf("Person is assignable to IAlive: %t\n", personType.AssignableTo(aliveInterface))
	fmt.Printf("Int is assignable to Temperature: %t\n", intType.AssignableTo(temperatureType))
	fmt.Printf("Int is convertible to Temperature: %t\n", intType.ConvertibleTo(temperatureType))
	fmt.Printf("Temperature is convertible to Int: %t\n", temperatureType.ConvertibleTo(intType))

	fmt.Println("--------------------------------------")

	// Создание нового объекта
	personType = TypeOf[Person]()
	personVal = reflect.New(personType)
	personVal.Elem().Field(0).SetString("Vovan")
	personVal.Elem().Field(1).SetInt(25)
	person = personVal.Elem().Interface().(Person)
	fmt.Println(person)

	fmt.Println("--------------------------------------")

	// Создание срезов
	sliceType = TypeOf[[]int]()
	intSliceValue := reflect.MakeSlice(sliceType, 0, 0)
	intSlice := ValueOf[[]int](intSliceValue)
	fmt.Println(intSlice)

	// Создание словарей
	mapType = TypeOf[map[string]int]()
	mapValue := reflect.MakeMap(mapType)
	myMap = ValueOf[map[string]int](mapValue)
	fmt.Println(myMap)

	fmt.Println("--------------------------------------")

	// Манипулирование срезами
	v := 10
	rv := reflect.ValueOf(v)
	intSliceValue = reflect.Append(intSliceValue, rv)
	intSlice = ValueOf[[]int](intSliceValue)
	fmt.Println(intSlice)

	// Манипулирование словарями
	k := "hello"
	rk := reflect.ValueOf(k)
	mapValue.SetMapIndex(rk, rv)
	myMap = ValueOf[map[string]int](mapValue)
	fmt.Println(myMap)

	fmt.Println("--------------------------------------")

	// Динамическое создание структур
	props := map[string]any{
		"Prop1": 1,
		"Prop2": "str",
		"Prop3": []int{10, 20, 30},
	}
	var structProps []reflect.StructField
	for propName, value := range props {
		propType := reflect.TypeOf(value)
		sf := reflect.StructField{
			Name: propName,
			Type: propType,
		}
		structProps = append(structProps, sf)
	}
	structType := reflect.StructOf(structProps)
	structObject := reflect.New(structType)
	for propName, value := range props {
		property := structObject.Elem().FieldByName(propName)
		property.Set(reflect.ValueOf(value))
	}
	fmt.Printf("%+v\n", structObject)

	fmt.Println("--------------------------------------")

	// Вызов функций через reflection
	function := Sum
	rf := reflect.TypeOf(function)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(function)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		out := vf.Call(in)
		fmt.Printf("calling %s...\n", runtime.FuncForPC(vf.Pointer()).Name())
		return out
	})
	var wSum = wrapperF.Interface().(func(int, int) int)
	fmt.Println(wSum(1, 2))
}

func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func ValueOf[T any](reflectValue reflect.Value) T {
	return reflectValue.Interface().(T)
}

func Sum(a, b int) int {
	return a + b
}
