package services

import (
	"SecondProject/db"
	"SecondProject/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const Authorization = `1.Sign in
2.Exit`


func Authorizationss(Db *sql.DB) {
	fmt.Println(Authorization)
	var k int
	_, err := fmt.Scan(&k)
	if err != nil {
		log.Println("Try again", err)
	}
	switch k {
	case 1:
		User := AuthorizationOperation(Db)
		if User.Admin == true {
			AdminOperation(Db, User)
		} else {
			UserOperation(Db, User)
		}
	case 2:
		os.Exit(2)
	default:
		log.Println("DEFAULT")
	}
}
func AuthorizationOperation(Db *sql.DB) (user models.User) {
	var login, password string
	fmt.Println("Login:")
	_, err := fmt.Scan(&login)
	if err != nil {
		log.Println("Incorrect password", err)
	}
	fmt.Println("Password:")
	_, err = fmt.Scan(&password)
	if err != nil {
		log.Println("Incorrect password", err)
	}
	row := Db.QueryRow(db.SelectUser, login, password)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Admin,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println("You have not sign up!!")
		os.Exit(1)
	}
	return
}