package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Ping struct {
	Name string `json:"name"`
}

type Pagination struct {
	Total int `json:"total"`
	List  []Ping
}

var pingSML []Ping
var resultPagination Pagination

func PingHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resultPagination = Pagination{
		List:  pingSML,
		Total: len(pingSML),
	}
	json.NewEncoder(w).Encode(resultPagination)
}

func CreatePingHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var ping Ping
	_ = json.NewDecoder(r.Body).Decode(&ping)
	pingSML = append(pingSML, ping)
	json.NewEncoder(w).Encode(ping)
}

func main() {
	pingSML = append(pingSML, Ping{Name: "Hello wordld"})
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandle).Methods("GET")
	r.HandleFunc("/ping", CreatePingHandle).Methods("POST")
	http.Handle("/", r)
	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
