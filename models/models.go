package models

type User struct {
	ID    int
	Name  string
	Email string
	// City  string
}

type Post struct {
	ID     int
	Title  string
	Body   string
	UserID int
	User   *User `pg:"rel:has-one"`
}
