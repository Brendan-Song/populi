package model

type User struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	Password  []byte
}
