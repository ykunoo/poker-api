package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// データベース接続情報
	connStr := "user=postgres password=postgres dbname=poker-db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open the DB connection: %v", err)
	}
	defer db.Close()

	// データベース接続確認
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Connected to the database!")

	// テーブル作成SQLクエリ
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS results (
		id SERIAL PRIMARY KEY,
		request_id VARCHAR(50) NOT NULL,
		hand TEXT NOT NULL,
		result TEXT NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// クエリを実行してテーブルを作成
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Table 'results' created successfully!")
}
