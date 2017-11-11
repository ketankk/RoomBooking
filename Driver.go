package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
	"time"
	"log"
	"encoding/json"
)
const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "bookit"
	DB_USER = "booking"
	DB_PASS = "booking"
)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
var people []Person


func main(){

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})


	router:=mux.NewRouter()
	router.HandleFunc("/people",GetPeople ).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000",router))

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func main1() {
	fmt.Println("Hello GoLang")

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT into booking(booked_room,booked_by,st_time,end_time,booked_on,meet_type,purpose_det) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	t:=time.Now()

	tf:=t.Format("2006-01-02 15:04:05")
	fmt.Print(tf)

	// "2017-11-10 20:15:10"
	_, err = stmt.Exec(1, 1, tf, tf, tf, "Scrum", "Daily standup Data Engineering Team")
	if err != nil {
		panic(err.Error())
	}

}
