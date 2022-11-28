package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"proyectoBD/src/config"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	dbUri         string
)

func init() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbUri = config.DBUri

	InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func dBInit() *sql.DB {
	var err error
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}

func QueryDB(query string) (*sql.Rows, error) {
	db := dBInit()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func InsertDB(insert string) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(insert)
	if err != nil {
		return 0, err
	}
	return d.LastInsertId()
}

func DeleteDB(delete string) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(delete)
	if err != nil {
		return 0, err
	}
	return d.RowsAffected()
}

func UpdateDB(update string) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(update)
	if err != nil {
		return 0, err
	}
	return d.RowsAffected()
}
