package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"json"
	"net/http"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

var noteStore = make(map[string]Note)

var id int

// GetNoteHandler - GET - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshall(notes)
	if err != nil { //This should not be in production
		panic(err)
	}
	w.WriteHeader(http.StatusOk)
	w.Write(j)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening....localhost:8080")
	server.ListenAndServe()
}
