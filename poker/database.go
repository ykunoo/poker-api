package poker

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL ドライバ
)

var db *sql.DB

// データベース接続の初期化
func InitDB() {
	var err error

	// データベース接続情報
	connStr := "user=youruser password=yourpassword dbname=yourdbname sslmode=disable"

	// データベース接続の確立
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	// 接続確認
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Successfully connected to the database")
}

// DBインスタンスの取得
func GetDB() *sql.DB {
	return db
}
