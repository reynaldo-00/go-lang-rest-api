package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Note struct {
	ID       string   `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	TextBody string   `json:"textBody,omitempty"`
	Tags     []string `json:"tags"`
}

var notes []Note

var noteId int = 2

func GetNotesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func GetNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range notes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Note{})
}

func CreateNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	var Note Note
	_ = json.NewDecoder(req.Body).Decode(&Note)
	Note.ID = strconv.Itoa(noteId)
	noteId = noteId + 1
	notes = append(notes, Note)
	json.NewEncoder(w).Encode(notes)
}

func DeleteNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range notes {
		if item.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(notes)
}

func EditNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var Note Note
	_ = json.NewDecoder(req.Body).Decode(&Note)
	for index, item := range notes {
		if item.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			notes = append(notes, Note)
			break
		}
	}
	json.NewEncoder(w).Encode(notes)

}

func main() {
	router := mux.NewRouter()

	notes = append(notes, Note{ID: "0", Title: "Note Title", TextBody: "Note Body", Tags: []string{"tag1", "tag2"}})
	notes = append(notes, Note{ID: "1", Title: "Note Title 2", TextBody: "Note Body 2"})

	router.HandleFunc("/note/get/all", GetNotesEndpoint).Methods("GET")
	router.HandleFunc("/note/get/{id}", GetNoteEndpoint).Methods("GET")
	router.HandleFunc("/note/create", CreateNoteEndpoint).Methods("POST")
	router.HandleFunc("/note/delete/{id}", DeleteNoteEndpoint).Methods("DELETE")
	router.HandleFunc("/note/edit/{id}", EditNoteEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
