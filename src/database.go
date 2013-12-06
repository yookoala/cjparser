package cjparser

/**
 * generic sqlite3 interface
 * to manage multiple import source into 1 single database file
 * queries are handled by individual SourceHandler implementation
 */

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DataEntry []string

type SourceHandler interface {
	Init(uni *ConvertDB) (err error)
	ParseLine(line string) (item DataEntry, err error)
	Insert(tx *sql.Tx, item DataEntry) (err error)
}

type ConvertDB struct {
	Filename string
	DB       *sql.DB
	conn     bool
	handlers map[string]SourceHandler
}

func (db *ConvertDB) InitDb() (err error) {
	return
}

func (db *ConvertDB) Register(name string, h SourceHandler) (err error) {
	if db.conn != true {
		err := db.Open()
		if err != nil {
			return err
		}
		db.handlers = make(map[string]SourceHandler)
	}
	db.handlers[name] = h
	return h.Init(db)
}

func (db *ConvertDB) Insert(name string, tx *sql.Tx, item []string) (err error) {
	if handler, ok := db.handlers[name]; ok {
		handler.Insert(tx, item)
		return nil
	}
	return errors.New(fmt.Sprintf("Handler \"%s\" not found", name))
}

func (db *ConvertDB) Open() (err error) {
	sdb, err := sql.Open("sqlite3", db.Filename)
	if err != nil {
		return
	}
	db.DB = sdb
	db.conn = true
	return nil
}

func (db *ConvertDB) Close() (err error) {
	defer db.DB.Close()
	_, err = db.DB.Exec("vacuum")
	if err != nil {
		panic(err)
	}
	return nil
}
