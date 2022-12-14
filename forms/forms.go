// forms.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

type Users struct {
	UserID       int    `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          string `json:"age"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
}

type Note struct {
	NoteID      int       `json:"note_id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Information string    `json:"information"`
	Time        time.Time `json:"time"`
	Status      string    `json:"status"`
	Delegation  string    `json:"delegation"`
	Users       string    `json:"users"`
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// do something with details
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		submit := r.FormValue("submit")

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

		if submit == "addUsers" {
			userDetails := Users{
				FirstName:    r.FormValue("FirstName"),
				LastName:     r.FormValue("LastName"),
				Age:          r.FormValue("Age"),
				EmailAddress: r.FormValue("EmailAddress"),
				PhoneNumber:  r.FormValue("PhoneNumber"),
			}

			//Adding user info to database
			sqlAddUsers := `INSERT INTO users (firstname, lastname, age, phonenumber, emailaddress)
							VALUES ($1, $2, $3, $4, $5)`
			_, err = db.Exec(sqlAddUsers, userDetails.FirstName, userDetails.LastName, userDetails.Age, userDetails.PhoneNumber, userDetails.EmailAddress)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("\nUser Inserted successfully!")
			}
			_ = userDetails
		} else if submit == "addNotes" {
			addNoteDetails := Note{
				UserID:      r.FormValue("UserID"),
				Name:        r.FormValue("noteName"),
				Information: r.FormValue("noteInformation"),
				Status:      r.FormValue("noteStatus"),
				Delegation:  r.FormValue("noteDelegation"),
				Users:       r.FormValue("userNote"),
			}

			var time = time.Now()

			//Adding note info to database
			sqlAddNotes := `INSERT INTO notes (userid, name, information, time, status, delegation, users)
							VALUES ($1, $2, $3, $4, $5, $6, $7)`
			_, err = db.Exec(sqlAddNotes, addNoteDetails.UserID, addNoteDetails.Name, addNoteDetails.Information, time, addNoteDetails.Status, addNoteDetails.Delegation, addNoteDetails.Users)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("\nNote Inserted successfully!")
			}
		}

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
