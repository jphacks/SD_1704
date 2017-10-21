package model

import (
	"database/sql"
)

// 全て値があるときに利用する
func InsertUser(db *sql.DB, nickname string, email string, password string) error {
	q := `insert into users (nickname, email, password) values ($1,$2,$3)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nickname, email, password)
	return err
}


