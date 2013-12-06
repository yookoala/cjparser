package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// hander for the Unihan_Variants.txt
type CangjieDataHandler struct {
}

func (h CangjieDataHandler) Init(db *ConvertDB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS Cangjie (
			id INTEGER PRIMARY KEY,
			unicode TEXT,
			character TEXT,
			version TEXT,
			category TEXT,
			code TEXT,
			serial TEXT
		)
	`
	_, err = db.DB.Exec(query)

	return
}

func (h CangjieDataHandler) Insert(tx *sql.Tx, item interface{}) (err error) {
	itemVal := item.(cangjieValue)
	stmt, err := tx.Prepare(`INSERT INTO Cangjie (
		unicode,
		character,
		version,
		category,
		code,
		serial
	) VALUES (
		?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(itemVal.Unicode, itemVal.Character, itemVal.Version,
		itemVal.Category, itemVal.Code, itemVal.Serial)
	return
}
