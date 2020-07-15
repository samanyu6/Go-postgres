package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name`
	Age  int32  `json:"age"`
}

type Result struct {
	Message string `json:"message"`
}

// Run main test function for all the middleware handlers.
func TestMain(m *testing.M) {
	fmt.Println(" TESTING MIDDLEWARE HANDLERS")
	os.Exit(m.Run())
	fmt.Println(" MIDDLEWARE COMPLETED ")
}

func TestCeateUser(t *testing.T) {

	user := &User{
		Name: "test",
		Age:  123,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Failed in marshalling json")
	}

	req, err := http.NewRequest("POST", "/api/user", bytes.NewBuffer(userJson))
	if err != nil {
		t.Errorf("Error making request, %v", err)
	}

	rec := httptest.NewRecorder()

	CreateUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Could not execute Create user with error %v", res.StatusCode)
	}

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading output %v", err)
	}

	var val Result
	err = json.Unmarshal(ret, &val)
	if err != nil {
		t.Errorf("error unmarshalling %v", err)
	}

	if string(val.Message) == "Error" {
		t.Errorf("error executing.")
	}
}

func TestGetUser(t *testing.T) {

	// Request url
	req, err := http.NewRequest("GET", "/api/usr", nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	rec := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	GetUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Failed with status %v", res.StatusCode)
	}

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading output %v", err)
	}

	var val Result
	err = json.Unmarshal(ret, &val)
	if err != nil {
		t.Errorf("error unmarshalling %v", err)
	}

	if string(val.Message) == "Error" {
		t.Errorf("error executing.")
	}
}

func TestUpdate(t *testing.T) {
	user := User{
		Name: "test_updated",
		Age:  34,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error marshalling user.")
	}

	req, err := http.NewRequest("POST", "/api/update", bytes.NewBuffer(userJson))
	if err != nil {
		t.Errorf("Error making new request.")
	}

	rec := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	UpdateUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Could not execute Create user with error %v", res.StatusCode)
	}

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading output %v", err)
	}

	var val Result
	err = json.Unmarshal(ret, &val)
	if err != nil {
		t.Errorf("error unmarshalling %v", err)
	}

	if string(val.Message) == "Error" {
		t.Errorf("error executing.")
	}

}

func TestDelete(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/api/delete", nil)
	if err != nil {
		t.Errorf("Error making request, %v", err)
	}

	rec := httptest.NewRecorder()

	vars := map[string]string{
		"Id": "1",
	}

	req = mux.SetURLVars(req, vars)

	DeleteUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Could not execute Create user with error %v", res.StatusCode)
	}

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading output %v", err)
	}

	var val Result
	err = json.Unmarshal(ret, &val)
	if err != nil {
		t.Errorf("error unmarshalling %v", err)
	}

	if string(val.Message) == "Error" {
		t.Errorf("error executing.")
	}
}
