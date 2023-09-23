package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Accounts struct {
	ID         int `json:"id"`
	Balance    float32 `json:"balance"`
	TargetGoal float64 `json:"targetGoal"`
	CreatedAt  time.Time `json:"createdAt"`
}
func CreateAccountTable(db *sql.DB) error {
	query := `
	DROP TABLE IF EXISTS accounts;
	CREATE TABLE accounts (
		id INT AUTO_INCREMENT,
		balance DECIMAL(10, 2) NOT NULL,
		targetgoal FLOAT NOT NULL,
		createdat DATETIME,
		PRIMARY KEY (id)
	);`
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
func CreateAccount(ac Accounts, db *sql.DB) (int64, error) {
	createdAt := time.Now()
	result, err := db.Exec("INSERT INTO users (balace, target_goal, created_at) VALUES(?, ?, ?)", ac.Balance, ac.TargetGoal, createdAt)
	if err != nil {
		return 0, fmt.Errorf( "createUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf( "createUser: %v", err)
	}
	return id, nil
}
func GetAllAccounts(name string, db *sql.DB) ([]Accounts, error) {
    // An Accountss slice to hold data from returned rows.
    var accounts []Accounts

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var ac Accounts
        if err := rows.Scan(&ac.ID, &ac.Balance, &ac.TargetGoal); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        accounts = append(accounts, ac)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return accounts, nil
}