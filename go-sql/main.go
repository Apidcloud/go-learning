package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ProductDto struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)
	insertedKey := insertProduct(db, ProductDto{"example2", 15.55, true})

	var name string
	var price float64
	var available bool

	query := `SELECT name, price, available FROM product WHERE id = $1`
	err = db.QueryRow(query, insertedKey).Scan(&name, &price, &available)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No rows found with ID %d", 1)
		}
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Price: %f\n", price)
	fmt.Printf("Available: %t\n", available)

	products := []ProductDto{}

	rows, err := db.Query("SELECT name, price, available FROM product")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var nameRead string
	var priceRead float64
	var availableRead bool

	for rows.Next() {
		err := rows.Scan(&nameRead, &priceRead, &availableRead)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, ProductDto{nameRead, priceRead, availableRead})
	}

	fmt.Println(products)
}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS Product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6, 2) NOT NULL,
		available BOOLEAN,
		createdAt timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, dto ProductDto) int {
	query := `INSERT INTO Product (name, price, available)
		VALUES ($1, $2, $3) RETURNING id`

	var primaryKey int
	err := db.QueryRow(query, dto.Name, dto.Price, dto.Available).Scan(&primaryKey)

	if err != nil {
		log.Fatal(err)
	}

	return primaryKey
}
