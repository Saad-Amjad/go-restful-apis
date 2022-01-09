package main

import (
	"database/sql"
	"go-restful-apis/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	db := initDB("storage.db")
	migrate(db)

	e.GET("/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetBooks(db))
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS books(
	  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  name VARCHAR NOT NULL
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
