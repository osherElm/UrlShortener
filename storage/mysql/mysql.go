package mysql

import (
	"UrlShortener/storage"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	db *sql.DB
}

//Initialize receives the parameters required to setup a mysql db
func Initialize(user string, password string, dbname string) (storage.Storage, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	msql := mysql{db: db}
	return msql, nil
}

func (m mysql) Save(url string) (string, error) {

	stmtIns, err := m.db.Prepare(`INSERT INTO urls (url, count, visited) VALUES(?,?,?)`)
	if err != nil {
		return "", err
	}
	_, err = stmtIns.Exec(url, 0, false)
	if err != nil {
		return "", fmt.Errorf("mysql: could not execute statement: %v", err)
	}

	defer stmtIns.Close()
	return url, nil
}

func (m mysql) Close() error {
	return m.db.Close()
}
