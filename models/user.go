package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// the id product of the user
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateUserTable(db *sql.DB) error {
	query := `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		firstname TEXT NOT NULL,
		lastname TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
func HashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}
func CreateUser(user User, db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO users (firstname, lastname, password, email) VALUES(?, ?, ?)", user.Firstname, user.Lastname, HashPassword(user.Password), user.Email)
	if err != nil {
		return 0, fmt.Errorf( "createUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf( "createUser: %v", err)
	}
	return id, nil
}