package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id      bson.ObjectId `bson:"_id"`
	Model   string        `bson:"model"`
	Company string        `bson:"company"`
	Price   int           `bson:"price"`
}

func main() {
	// Подключение к БД
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// получаем коллекцию products в базе данных productdb
	productCollection := session.DB("productdb").C("products")

	fmt.Println("---------------------------------")

	// Вставка данных
	p1 := &Product{Id: bson.NewObjectId(), Model: "iPhone 8", Company: "Apple", Price: 64567}
	// добавляем один объект
	err = productCollection.Insert(p1)
	if err != nil {
		fmt.Println(err)
	}

	p2 := &Product{Id: bson.NewObjectId(), Model: "Pixel 2", Company: "Google", Price: 58000}
	p3 := &Product{Id: bson.NewObjectId(), Model: "Xplay7", Company: "Vivo", Price: 49560}
	// добавляем два объекта
	err = productCollection.Insert(p2, p3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------------------------------")

	// Выборка данных
	query := bson.M{
		"price": bson.M{
			"$gt": 50000,
		},
	}
	products := []Product{}
	productCollection.Find(query).All(&products)
	// productCollection.Find(query).One(&product)

	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("---------------------------------")

	// Обновление данных
	updateInfo, err := productCollection.UpdateAll(bson.M{"model": "iPhone 8"}, bson.M{"$set": bson.M{"price": 45000}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Updated %d items\n", updateInfo.Updated)
	}
	PrintAllProducts(productCollection)

	fmt.Println("---------------------------------")

	// Удаление данных
	deleteInfo, err := productCollection.RemoveAll(bson.M{"company": "Vivo"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Deleted %d items\n", deleteInfo.Removed)
	}
	PrintAllProducts(productCollection)
}

func PrintAllProducts(productCollection *mgo.Collection) {
	products := []Product{}
	productCollection.Find(bson.M{}).All(&products)
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}
}
