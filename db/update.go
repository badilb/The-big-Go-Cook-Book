package db

import (
	"database/sql"
	"fmt"
)

// Обновление
func Update(db *sql.DB) {
	// обновляем строку с id=1
	result, err := db.Exec("update Products set price = $1 where id = $2", 69000, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}
