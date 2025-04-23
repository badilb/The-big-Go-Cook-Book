package db

import (
	"database/sql"
	"fmt"
)

func AddProduct(db *sql.DB) {
	// Простое добавление данных без возврата ID
	result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Добавлено строк: %d\n", rowsAffected)

	// Добавление, если нужно получить id добавленного обьекта
	var id int
	db.QueryRow("insert into Products (model, company, price) values ('Mate 10 Pro', $1, $2) returning id",
		"Huawei", 35000).Scan(&id)
	fmt.Println(id)

}
