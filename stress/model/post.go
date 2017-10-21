package model

import "database/sql"

func InsertPost(db *sql.DB, description string, userId int64) error {
	q := `insert into posts (description, user_id) values ($1,$2)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(description, userId)
	return err
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
