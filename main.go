package main

import (
	"database/sql"

	mydb "github.com/brandaogabriel7/full-cycle-arquitetura-hexagonal/adapters/db"
	"github.com/brandaogabriel7/full-cycle-arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := mydb.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)

}
