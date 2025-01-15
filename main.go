package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	db := NewPostgresClient()
	if db == nil {
		log.Fatalf("Failed to initialize PostgreSQL client")
	}

	filename := "path/to/your/largefile.csv"
	tableName := "your_table"
	columns := "(column1, column2, column3)" // Cambia esto según la estructura de tu tabla

	err := copyCSVToPostgres(db, filename, tableName, columns)
	if err != nil {
		log.Fatalf("Error copying CSV to PostgreSQL: %v", err)
	}

	log.Println("Records successfully saved to PostgreSQL!")
}

func copyCSVToPostgres(db *sql.DB, filename, tableName, columns string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	copyQuery := fmt.Sprintf("COPY %s %s FROM STDIN WITH CSV HEADER", tableName, columns)
	_, err = db.Exec(copyQuery)
	if err != nil {
		return fmt.Errorf("could not prepare COPY statement: %v", err)
	}

	return nil
}
