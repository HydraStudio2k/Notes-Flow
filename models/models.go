package models

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	Uuid     string
}

type Note struct {
	Id     int
	Author string
	Status string
	Note   string
	Uuid   string
}
