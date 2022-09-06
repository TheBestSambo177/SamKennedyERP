package main

//Convert JSON to Go struct
//https://mholt.github.io/json-to-go/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq" // Interface to PostgreSQL library
)

//Postgres connecting
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "erp"
)

type Results struct {
	Marks    []Mark    `json:"marks"`
	Students []Student `json:"students"`
}
type Mark struct {
	StudentID int     `json:"student_id"`
	Class     string  `json:"class"`
	Mark      float64 `json:"mark"`
}
type Student struct {
	StudentID   int    `json:"student_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	PhoneNumebr string `json:"phone_numebr"`
	Suburb      string `json:"suburb"`
	City        string `json:"city"`
}

type listOfNotes struct {
	Notes []Note `json:"notes"`
	Users []User `json:"users"`
}

type Note struct {
	NoteID      int       `json:"note_id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Information string    `json:"information"`
	Time        time.Time `json:"time"`
	Status      string    `json:"status"`
	Delegation  string    `json:"delegation"`
	Users       string    `json:"users"`
}

type User struct {
	UserID       int    `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
}

// Go custom sorting
type ByWord []Student

func (s ByWord) Len() int {
	return len(s)
}

func (s ByWord) Swap(i, j int) {
	s[i].LastName, s[j].LastName = s[j].LastName, s[i].LastName
}

func (s ByWord) Less(i, j int) bool {
	return len(s[i].LastName) < len(s[j].LastName)
}

// partial array dumper
func dump(arr []Note) {
	fmt.Println("dump Notes")
	for i, v := range arr {
		fmt.Println("\nUSERID:", v.UserID, "\nNote Name:", v.Name, "\nNote Info:", v.Information, "\nNote Time:", v.Time, "\nNote Status:", v.Status, "\nNote Delegation:", v.Delegation, "\nNote Users:", v.Users)
		if i > 5 {
			break
		}
	}
}

//Test

// -----------------------------------------------------------------
func main() {

	// Create a string that will be used to make a connection later
	// Note Password has been left out, which is best to avoid issues when using null password
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

	fmt.Println("Connected successfully")

	sqlUser := `SELECT * FROM users LIMIT 100`
	sqlNotes := `SELECT * FROM notes LIMIT 100`

	userRows, err := db.Query(sqlUser) // $1 and $2 set here. Note sqlStatement could be replaced with literal string
	if err != nil {
		log.Fatal(err)
		fmt.Println("An error occurred when querying data!")
	}
	defer userRows.Close()

	noteRows, err := db.Query(sqlNotes) // $1 and $2 set here. Note sqlStatement could be replaced with literal string
	if err != nil {
		log.Fatal(err)
		fmt.Println("An error occurred when querying data!")
	}
	defer noteRows.Close()

	for userRows.Next() {

		var UserID int
		var FirstName string
		var LastName string
		var Age int
		var PhoneNumber string
		var EmailAddress string

		switch err = userRows.Scan(&UserID, &FirstName, &LastName, &Age, &PhoneNumber, &EmailAddress); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(UserID, "|", FirstName, "|", LastName, "|", Age, "|", PhoneNumber, "|", EmailAddress)
		default:
			fmt.Println("SQL query error occurred: ")
			panic(err)
		}

	}

	//get any error encountered during User Test
	err = userRows.Err()
	if err != nil {
		panic(err)

	}

	for noteRows.Next() {

		var NoteID int
		var UserID int
		var Name string
		var Information string
		var Time string
		var Status string
		var Delegation string
		var Users string

		switch err = noteRows.Scan(&NoteID, &UserID, &Name, &Information, &Time, &Status, &Delegation, &Users); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(NoteID, "|", UserID, "|", Name, "|", Information, "|", Time, "|", Status, "|", Delegation, "|", Users)
		default:
			fmt.Println("SQL query error occurred: ")
			panic(err)
		}

	}

	//get any error encountered during User Test
	err = noteRows.Err()
	if err != nil {
		panic(err)

	}

	//Json file stuff
	var (
		results listOfNotes
	)

	data, err := ioutil.ReadFile("./notes.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &results)
	fmt.Println("Unsorted notes")
	dump(results.Notes)

}
