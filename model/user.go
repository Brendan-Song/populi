package model

type user struct {
    ID		int64
    Username	string
    FirstName	string
    LastName	string
    Password	[]byte
}