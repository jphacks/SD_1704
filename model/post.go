package model

import "database/sql"

func InsertPost(db *sql.DB, description string, userId int64) (int64, error) {
	q := `insert into posts (description, user_id) values ($1,$2)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(description, userId)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func GetPostsByUserId(db *sql.DB, userId int64) ([]Post, error) {
	q := `select * from posts where user_id=$1`

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	return ScanPosts(rows)
}

func GetPostById(db *sql.DB, Id int64) (Post, error) {
	q := `select * from posts where id=$1`

	var post Post

	stmt, err := db.Prepare(q)
	if err != nil {
		return post, err
	}

	err = stmt.QueryRow(Id).Scan(&post.ID, &post.Description, &post.UserId, &post.Created)
	return post, err
}

func GetRandomPost(db *sql.DB) (Post, error) {
	q := `SELECT * FROM posts ORDER BY random() LIMIT 1`

	var post Post

	stmt, err := db.Prepare(q)
	if err != nil {
		return post, err
	}

	err = stmt.QueryRow().Scan(&post.ID, &post.Description, &post.UserId, &post.Created)
	return post, err
}
