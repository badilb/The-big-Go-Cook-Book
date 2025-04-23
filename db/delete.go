package db

import (
	"database/sql"
	"fmt"
)

// Удаление
func Delete(db *sql.DB) {
	// удаляем строку с id=2
	result, err := db.Exec("delete from Products where id = $1", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}
