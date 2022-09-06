package main

//Convert JSON to Go struct
//https://mholt.github.io/json-to-go/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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

// rng=3 is about 1ms
func delay(rng int) {
	for i := 0; i < rng; i++ {
		for j := 0; j < 1000000; j++ {
		}
	}
}

// binary search
func binarySearch(arr []Student, inT string) int {
	low := 0
	high := len(arr) - 1
	T := strings.ToLower(inT)

	for low <= high {
		var mid = low + (high-low)/2                      //middle of the list
		var midvalue = strings.ToLower(arr[mid].LastName) //get item to match with T

		switch {
		case midvalue == T:
			return mid
		case midvalue < T:
			low = mid + 1
		case midvalue > T:
			high = mid - 1
		}

	}

	return -1
}

func merge(fp []Student, sp []Student) []Student {
	var n = make([]Student, len(fp)+len(sp))

	var fpIndex = 0
	var spIndex = 0
	var nIndex = 0

	for fpIndex < len(fp) && spIndex < len(sp) {
		if fp[fpIndex].LastName < sp[spIndex].LastName {
			n[nIndex] = fp[fpIndex]
			fpIndex++
		} else if sp[spIndex].LastName < fp[fpIndex].LastName {
			n[nIndex] = sp[spIndex]
			spIndex++
		}
		nIndex++
	}

	for fpIndex < len(fp) {
		n[nIndex] = fp[fpIndex]
		fpIndex++
		nIndex++
	}

	for spIndex < len(sp) {
		n[nIndex] = sp[spIndex]
		spIndex++
		nIndex++
	}

	return n
}

func mergeSort(arr []Student) []Student {
	if len(arr) == 1 {
		return arr
	}

	var fp = mergeSort(arr[0 : len(arr)/2])
	var sp = mergeSort(arr[len(arr)/2:])
	delay(1)
	return merge(fp, sp)

}

func bubbleSort(arr []Student) []Student {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].LastName > arr[j+1].LastName {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
			delay(1)
		}

	}
	return arr
}

func selectionSort(arr []Student) []Student {

	for i := 0; i < len(arr)-1; i++ {
		var j = i + 1

		var minIndex = i

		if j < len(arr) {
			if arr[j].LastName < arr[minIndex].LastName {
				minIndex = j
			}
			j++
			delay(1)
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

	}
	return arr
}

func insertionSort(arr []Student) []Student {
	for i := 1; i < len(arr); i++ {
		key := arr[i].LastName
		keyStruct := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j].LastName {
			arr[j+1] = arr[j]
			j -= 1
			delay(1)
		}
		arr[j+1] = keyStruct
	}
	return arr
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
