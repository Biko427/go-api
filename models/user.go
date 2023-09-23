package models

type User struct {
	// the id product of the user
	ID int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

