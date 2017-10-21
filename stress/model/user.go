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

func GetUserByUserId(db *sql.DB, Id int64) (User, error) {
	q := `select * from users where id=$1`

	var user User
	stmt, err := db.Prepare(q)
	if err != nil {
		return user, err
	}

	err = stmt.QueryRow(Id).Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.Created)
	return user, err
}
