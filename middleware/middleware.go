package middleware

import (
	"encoding/json"
	"factly/database"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

// Create a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Sanitize response values
	var users database.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		log.Fatal("Unable to receive requested data")
	}

	// Open database connection listener
	db := database.CreateConn()
	defer db.Close()

	// Query string
	pQuery := `INSERT INTO users (name, age) VALUES ($1,$2) RETURNING Id`
	var id int64

	// Execute query
	err = db.QueryRow(pQuery, users.Name, users.Age).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Return message
	fmt.Printf("Inserted a single record %v", id)
	msg := fmt.Sprintf("Added data successfully %d", id)
	res := Response{
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Response data in map[string]string format and extract id from response
	resp := mux.Vars(r)
	id, err := strconv.Atoi(resp["Id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int for %v", err)
	}

	db := database.CreateConn()
	defer db.Close()

	// Delete query executing
	delQuery := `DELETE FROM users WHERE Id=$1`
	res, err := db.Exec(delQuery, id)
	if err != nil {
		log.Fatalf("Unable to delete id = %d, and err %v", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error getting rows affected %v", err)
	}

	// return message
	fmt.Println("Rows affected %v", rowsAffected)
	retMsg := fmt.Sprintf("Successful, records affected %v", rowsAffected)
	returnRes := Response{
		Message: retMsg,
	}

	json.NewEncoder(w).Encode(returnRes)
}
