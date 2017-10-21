package model

import "time"

type User struct {
	ID       int64
	Nickname string
	Email    string
	Password string
	Created  *time.Time
}

type Post struct {
	ID          int64
	Description string
	UserId      string
	Created     *time.Time
}

type Notification struct {
	ID      int64
	Time    *time.Time
	Set     bool
	Created *time.Time
}
