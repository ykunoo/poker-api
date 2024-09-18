package poker

import (
	"database/sql"
	"strings"
	"time"
)

// データベースに結果を挿入する関数
func insertResultTx(tx *sql.Tx, requestId string, hand string, result string) error {
	query := `INSERT INTO results (request_id, hand, result, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := tx.Exec(query, requestId, hand, result, time.Now())
	return err
}

// 評価結果とエラーをDBに保存する関数
func SaveResultsToDB(results []HandResult, errors []HandError) error {

	// データベース接続を取得
	db := GetDB()

	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // エラー発生時にロールバック
		}
	}()

	// 正常結果をDBに保存
	for _, result := range results {
		err = insertResultTx(tx, result.RequestId, result.Hand, result.Yaku)
		if err != nil {
			tx.Rollback() // エラー発生時にロールバック
			return err
		}
	}

	// エラー結果をDBに保存
	for _, handError := range errors {
		errorMessage := strings.Join(handError.ErrorMessages, ",")
		err = insertResultTx(tx, handError.RequestId, handError.Hand, errorMessage)
		if err != nil {
			tx.Rollback() // エラー発生時にロールバック
			return err
		}
	}

	// コミット
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
