package model

import (
	"database/sql"
	"fmt"
)

func CreateSession(db *sql.DB, sessionId string, userId int64) error {
	q := "insert into sessions (session_id, user_id) values (?,?)"

	result, err := db.Exec(q, sessionId, userId)
	if err != nil {
		fmt.Println(result)
	}
	return err
}

func SessionExistsByUserId(db *sql.DB, userId int64) (bool, error) {
	q := "select count(*) from sessions where user_id=?"

	var count int64
	if err := db.QueryRow(q, userId).Scan(&count); err != nil {
		return false, err
	}

	return count == 1, nil
}

func UpdateSessionByUserId(db *sql.DB, uuid string, userId int64) error {
	q := "update sessions set session_id=? where user_id =?"
	_, err := db.Exec(q, uuid, userId)
	return err
}

func GetUserIdByUuid(db *sql.DB, uuid string) (*int64, error) {
	q := "select user_id from sessions where session_id=?"

	var user_id int64
	if err := db.QueryRow(q, uuid).Scan(&user_id); err != nil {
		return nil, err
	}
	return &user_id, nil
}
