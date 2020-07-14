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

// Delete a user
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

// Get all users
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Connect db
	db := database.CreateConn()
	defer db.Close()

	// Array of users to return
	var users []database.User

	// Get all users query
	allSql := `SELECT * FROM users`
	rows, err := db.Query(allSql)
	if err != nil {
		log.Fatalf("Error executing get all users query : %v", err)
	}
	defer rows.Close()

	// Append all users to an array to return
	for rows.Next() {
		var u database.User

		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			log.Fatalf("Error scanning %v", err)
		}

		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

// Get specific user
func GetUser(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// retrieve id
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatalf("Error retrieving id from link %v", err)
	}

	// Create database connection
	db := database.CreateConn()
	defer db.Close()

	var user database.User

	// Sql query
	usrSql := `SELECT * FROM users WHERE Id=$1`
	row := db.QueryRow(usrSql, id)

	// Return user data
	msg := ""
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		msg = fmt.Sprintf("%v", err)
	} else {
		msg = fmt.Sprintf("Results %v", row)
	}

	resp := Response{
		Message: msg,
	}

	json.NewEncoder(w).Encode(resp)
}

// Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Extract id
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Error converting id type %v", err)
	}

	var user database.User

	// Decode request from user to change
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error decoding %v", err)
	}

	// Connect db
	db := database.CreateConn()
	defer db.Close()

	// SQL query
	updateSql := `UPDATE users SET Name=$1, Age=$2 WHERE Id=$3`
	res, err := db.Exec(updateSql, user.Name, user.Age, id)
	if err != nil {
		log.Fatal("Error updating table %v", err)
	}

	// retrieve updated results
	rowsAffect, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error retreiving rows affected %v", err)
	}

	fmt.Printf("Rows affected %v", rowsAffect)
	msg := fmt.Sprintf("Rows updated %v", rowsAffect)
	resp := Response{
		Message: msg,
	}

	json.NewEncoder(w).Encode(resp)
}
