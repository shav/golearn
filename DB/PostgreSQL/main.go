package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	// Подключение к БД
	connStr := ""
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.Ping()

	fmt.Println("---------------------------------")

	// Вставка данных
	result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	var id int
	db.QueryRow("insert into Products (model, company, price) values ('Mate 10 Pro', $1, $2) returning id",
		"Huawei", 35000).Scan(&id)
	fmt.Println("LastInsertId: ", id)

	fmt.Println("---------------------------------")

	// Выборка данных
	rows, err := db.Query("select * from Products where price > $1", 70000)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []product{}
	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("---------------------------------")

	// Обновление данных
	result, err = db.Exec("update Products set price = $1 where id = $2", 69000, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк

	// Выборка только одной строки
	row := db.QueryRow("select * from Products where id = $1", 3)
	prod := product{}
	err = row.Scan(&prod.id, &prod.model, &prod.company, &prod.price)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", prod)

	fmt.Println("---------------------------------")

	// Удаление данных
	result, err = db.Exec("delete from Products where id = $1", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}
