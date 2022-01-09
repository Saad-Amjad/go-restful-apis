package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BookCollection struct {
	Book []Book `json:"data"`
}

func GetBooks(db *sql.DB) BookCollection {
	sql := "SELECT * FROM books"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := BookCollection{}
	for rows.Next() {
		book := Book{}
		err2 := rows.Scan(&book.ID, &book.Name)
		if err2 != nil {
			panic(err2)
		}
		result.Book = append(result.Book, book)
	}
	return result
}
