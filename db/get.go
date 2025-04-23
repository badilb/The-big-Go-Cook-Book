package db

import (
	"database/sql"
	"fmt"
)

// структура product, которая соответствует данным в таблице Products
// нужна для работы с данными
type product struct {
	id      int
	model   string
	company string
	price   int
}

// Получение данных
func GetData(db *sql.DB) {
	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	// здесь мы с помощью rows.Next() перебираем каждую строку, полученную из базы данных, и отображаем все строки
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
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}

func GetDataWithID(db *sql.DB) {
	// Если надо получить только одну строку, то можно использовать метод QueryRow():
	row := db.QueryRow("select * from Products where id = $1", 2)
	prod := product{}
	err := row.Scan(&prod.id, &prod.model, &prod.company, &prod.price)
	if err != nil {
		panic(err)
	}
	fmt.Println(prod.id, prod.model, prod.company, prod.price)
}
