package models

type User struct {
	ID int64
	Name string
	Surname string
	Age int64
	Gender string
	Admin bool
	Login string
	Password string
	Remove bool
}