package main

//Convert JSON to Go struct
//https://mholt.github.io/json-to-go/

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

//Select all users
func selectUsers() {
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

	sqlUser := `SELECT * FROM users LIMIT 100`

	userRows, err := db.Query(sqlUser)
	if err != nil {
		log.Fatal(err)
		fmt.Println("An error occurred when querying data!")
	}
	defer userRows.Close()

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
}

//Select all notes
func selectNotes() {
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

	sqlNotes := `SELECT * FROM notes LIMIT 100`

	noteRows, err := db.Query(sqlNotes)
	if err != nil {
		log.Fatal(err)
		fmt.Println("An error occurred when querying data!")
	}
	defer noteRows.Close()

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
}

//Insert Users
func addUsers() {
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

	//Getting user info
	//Variables for user

	//Test
	//var fName string
	//var lName string
	//var age int
	//var phonenumber int
	//var emailaddress string

	//Reader to keep input on one line
	consoleReader := bufio.NewReader(os.Stdin)

	//Scanning user input
	fmt.Print("First Name: ")
	fName, _ := consoleReader.ReadString('\n')

	fmt.Print("Last Name: ")
	lName, _ := consoleReader.ReadString('\n')

	fmt.Print("Age: ")
	age, _ := consoleReader.ReadString('\n')

	fmt.Print("Phone Number: ")
	phonenumber, _ := consoleReader.ReadString('\n')

	fmt.Print("Email Address: ")
	emailaddress, _ := consoleReader.ReadString('\n')

	//Adding user info to database
	sqlAddUsers := `INSERT INTO users (firstname, lastname, age, phonenumber, emailaddress)
	VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(sqlAddUsers, fName, lName, age, phonenumber, emailaddress)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUser Inserted successfully!")
	}

}

//Insert Users
func addNotes() {
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

	//Getting user info
	//Variables for user
	var name string
	var information string
	var status string
	var delegation string
	var users string
	var time = time.Now()

	//Scanning user input
	fmt.Print("Name of note: ")
	fmt.Scanln(&name)

	fmt.Print("Information of note: ")
	fmt.Scanln(&information)

	fmt.Print("Status of note: ")
	fmt.Scanln(&status)

	fmt.Print("Delegation of note: ")
	fmt.Scanln(&delegation)

	fmt.Print("Users for note: ")
	fmt.Scanln(&users)

	//Adding note info to database
	sqlAddNotes := `INSERT INTO notes (userid, name, information, time, status, delegation, users)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(sqlAddNotes, 3, name, "Make burgers", time, "Doing", "Sam", "Sam")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nNote Inserted successfully!")
	}

}

//Remove Users
func removeUsers() {
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

	//Asks user for id to remove
	var userRemId int
	fmt.Println("What ID do you want to remove: ")
	fmt.Scanln(&userRemId)

	//Remove User from note table
	sqlRemUserN := `
	DELETE FROM notes
	WHERE UserID = $1;`
	res1, err := db.Exec(sqlRemUserN, userRemId)
	if err != nil {
		panic(err)
	}
	_, err = res1.RowsAffected()
	if err != nil {
		panic(err)
	}
	//Remove User from user table
	sqlRemUser := `
	DELETE FROM users
	WHERE UserID = $1;`
	res2, err := db.Exec(sqlRemUser, userRemId)
	if err != nil {
		panic(err)
	}
	_, err = res2.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("row deleted")

}

//Remove Notes
func removeNotes() {
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

	//Asks user for id to remove
	var noteRemId int
	fmt.Println("What ID do you want to remove: ")
	fmt.Scanln(&noteRemId)

	//Remove User from note table
	sqlRemNote := `
	DELETE FROM notes
	WHERE NoteID = $1;`
	res1, err := db.Exec(sqlRemNote, noteRemId)
	if err != nil {
		panic(err)
	} else {
		fmt.Print("Rows Affected: ")
		fmt.Println(res1.RowsAffected())
	}

}

//Search Users
func searchUsers() {
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

	//Select * from notes Where name = $1 or text = $1 or status = $1 or delegation = $1 or users = $1;

	var userSearchInput string
	fmt.Println("What do you want to search? ")
	fmt.Scanln(&userSearchInput)
	userSearch := userSearchInput + "%"

	//Search Users from user table
	sqlSearchUsers := `Select * from users Where firstname ILIKE $1 or lastname ILIKE $1;`

	searchUsers, err := db.Query(sqlSearchUsers, userSearch)
	if err != nil {
		log.Fatal(err)
		fmt.Println("An error occurred when querying data!")
	}
	defer searchUsers.Close()

	for searchUsers.Next() {

		var UserID int
		var FirstName string
		var LastName string
		var Age int
		var PhoneNumber string
		var EmailAddress string

		switch err = searchUsers.Scan(&UserID, &FirstName, &LastName, &Age, &PhoneNumber, &EmailAddress); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(UserID, "|", FirstName, "|", LastName, "|", Age, "|", PhoneNumber, "|", EmailAddress)
		default:
			fmt.Println("SQL query error occurred: ")
			panic(err)
		}

	}

}

//Search Notes
func searchNotes() {
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

	var userSearchInput string
	fmt.Println("What do you want to search? ")
	fmt.Scanln(&userSearchInput)
	userSearch := userSearchInput + "%"

	//Search Notes from Note table
	sqlSearchNotes := `Select * from notes Where name ILIKE $1 or information ILIKE $1 or status ILIKE $1 or delegation ILIKE $1 or users ILIKE $1;`

	searchNotes, err := db.Query(sqlSearchNotes, userSearch)
	if err != nil {
		log.Fatal(err)
		fmt.Println("An error occurred when querying data!")
	}
	defer searchNotes.Close()

	for searchNotes.Next() {

		var NoteID int
		var UserID int
		var Name string
		var Information string
		var Time string
		var Status string
		var Delegation string
		var Users string

		switch err = searchNotes.Scan(&NoteID, &UserID, &Name, &Information, &Time, &Status, &Delegation, &Users); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(NoteID, "|", UserID, "|", Name, "|", Information, "|", Time, "|", Status, "|", Delegation, "|", Users)
		default:
			fmt.Println("SQL query error occurred: ")
			panic(err)
		}

	}

}

func updateUsers() {
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

	var userID int
	fmt.Println("What userID do you want to change? ")
	fmt.Scanln(&userID)

	var userSection string
	fmt.Println("What section do you want to change? ")
	fmt.Scanln(&userSection)

	updateUserStatement := `UPDATE users set FirstName = $2, LastName = $3 Where userID = $1;`

	_, err = db.Exec(updateUserStatement, userID, "NewFirst", "NewLast")
	if err != nil {
		panic(err)
	}
}

//Global for var for user
var currentUserID int = 0

// -----------------------------------------------------------------
func main() {
	i := 1
	for i == 1 {
		//Shows the user the ID they are logged into. If not are say no user is logged in.
		if currentUserID != 0 {
			println("You are currently logged in as user", currentUserID)
		} else {
			println("No user is logged in right now.")
		}
		fmt.Println("Sign In (1) | Users (2) | Notes (3) | Sign Out (4) | End (5): ")
		var userOption int
		fmt.Scanln(&userOption)
		if userOption == 1 {
			fmt.Print("Enter ID: ")
			var signInInput int
			fmt.Scanln(&signInInput)
			currentUserID = signInInput
		} else if userOption == 5 {
			break
		} else if currentUserID == 0 {
			println("Please Sign to get full use of program")
		} else {
			if userOption == 5 {
				break
			} else if userOption == 2 {
				var userOptionUser string
				fmt.Println("Users: Select All (a) | Insert (i) | Remove (r) | Search (s) | Update (u) | Back (b): ")
				fmt.Scanln(&userOptionUser)
				if userOptionUser == "a" || userOptionUser == "A" {
					selectUsers()
				} else if userOptionUser == "i" || userOptionUser == "I" {
					addUsers()
				} else if userOptionUser == "r" || userOptionUser == "U" {
					removeUsers()
				} else if userOptionUser == "s" || userOptionUser == "S" {
					searchUsers()
				} else if userOptionUser == "u" || userOptionUser == "U" {
					updateUsers()
				} else if userOptionUser == "b" || userOptionUser == "B" {
					fmt.Println("Going Back")
				} else {
					fmt.Println("Not a option")
				}

			} else if userOption == 3 {
				var userOptionNote string
				fmt.Println("Notes: Select All (a) | Insert (i) | Remove (r) | Search (s) | Back (b): ")
				fmt.Scanln(&userOptionNote)
				if userOptionNote == "a" || userOptionNote == "A" {
					selectNotes()
				} else if userOptionNote == "i" || userOptionNote == "I" {
					addNotes()
				} else if userOptionNote == "r" || userOptionNote == "R" {
					removeNotes()
				} else if userOptionNote == "s" || userOptionNote == "S" {
					searchNotes()
				} else if userOptionNote == "b" || userOptionNote == "B" {
					fmt.Println("Going Back")
				} else {
					fmt.Println("Not a option")
				}
			} else if userOption == 4 {
				currentUserID = 0
				fmt.Println("You have been signed out.")
			}
		}
	}

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
