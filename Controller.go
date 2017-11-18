package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
	router:=Router()
	Rooms(*router)
	Users(*router)

	router.HandleFunc("/rooms", GetRoom).Methods("GET")
log.Print("Rooms API")
	http.ListenAndServe(":1234", router);


}

func Router() *mux.Router {
	router := mux.NewRouter()
	return router

}
func Rooms(router mux.Router) mux.Router{
	router.HandleFunc("/rooms", GetRoom).Methods("GET")
	router.HandleFunc("/rooms", PutRoom).Methods("PUT")
	router.HandleFunc("/rooms", DeleteRoom).Methods("DELETE")

	return router

}
func Users(router mux.Router) {
	router.HandleFunc("/users", GetUser).Methods("GET")


}
func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Users")

}

func DeleteRoom(writer http.ResponseWriter, r *http.Request) {

}
func PutRoom(writer http.ResponseWriter, r *http.Request) {

}

func GetRoom(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("Rooms")
}
