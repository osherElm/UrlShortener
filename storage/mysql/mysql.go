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

func Initialize(user string, password string, dbname string) (storage.Storage, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// proper error handling instead of panic in your app
		return nil, err
	}

	msql := mysql{db: db}
	return msql, nil
}

func (m mysql) Save(url string) (string, error) {
	return url, nil
}

func (m mysql) Close() error {
	return m.db.Close()
}
