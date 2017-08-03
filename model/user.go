package model

type User struct {
	ID        int64  `json:"ID"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Password  []byte `json:"Password"`
}
