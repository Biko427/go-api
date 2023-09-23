package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Biko427/go-api/models"
)

func SaveUser(rw http.ResponseWriter, r *http.Request) {
	
}

func SaveAccount(rw http.ResponseWriter, r *http.Request) {
	var account models.Accounts
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&account)
	}
}