package database

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println(" TESTING DATABASE")
	os.Exit(m.Run())
	fmt.Println(" DONE TESTING DATABASE")
}

func TestDbConn(t *testing.T) {

	db := CreateConn()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Errorf("No database connection.")
	}
}
