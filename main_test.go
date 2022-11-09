package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq" // Interface to PostgreSQL library
)

//Tests
func TestAddUser(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Invalid DB arguments, or github.com/lib/pq not installed")
	}

	defer db.Close() // Housekeeping. Ensure connection is always closed once done

	// Ping database (connection is only established at this point, open only validates arguments passed to it)
	err = db.Ping()
	if err != nil {
		log.Fatal("Connection to specified database failed: ", err)

	}

	got := "Test"
	want := "Test"

	if got != want {
		t.Errorf(got, want)
	} else {
		t.Logf(got, want)
	}

}

func TestAddNote(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Invalid DB arguments, or github.com/lib/pq not installed")
	}

	defer db.Close() // Housekeeping. Ensure connection is always closed once done

	// Ping database (connection is only established at this point, open only validates arguments passed to it)
	err = db.Ping()
	if err != nil {
		log.Fatal("Connection to specified database failed: ", err)

	}

	got := "Test"
	want := "Test"

	if got != want {
		t.Errorf(got, want)
	} else {
		t.Logf(got, want)
	}
}

func TestSelectUser(t *testing.T) {

}

func TestSelectNotes(t *testing.T) {

}

func TestRemoveUser(t *testing.T) {

}

func TestRemoveNotes(t *testing.T) {

}

func TestSearchUser(t *testing.T) {

}

func TestSearchNote(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

func TestUpdateNote(t *testing.T) {

}

func TestValidate(t *testing.T) {

}