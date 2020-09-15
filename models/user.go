package models

type User struct {
	Name string
	Id   int
}

var AllUsers = make(map[int]User)
