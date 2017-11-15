package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

func main() {
	router:=mux.NewRouter()
	router.HandleFunc("/rooms",GetRoom).Methods("GET")
	router.HandleFunc("/rooms",PutRoom).Methods("PUT")
	router.HandleFunc("/rooms",DeleteRoom).Methods("DELETE")

	http.ListenAndServe(":1234",router);
}
func DeleteRoom(writer http.ResponseWriter, r *http.Request) {

}
func PutRoom(writer http.ResponseWriter, r *http.Request) {

}

func GetRoom(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("people")
}