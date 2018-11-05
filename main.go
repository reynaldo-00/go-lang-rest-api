package main

import (
	"encoding/json"
	"fmt"
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

type Error struct {
	Message string `json:"message,omitempty"`
}

var notes []Note

var noteId int = 2

func GetNotesEndpoint(w http.ResponseWriter, req *http.Request) {
	// json.NewEncoder(w).Encode(notes)

	if len(notes) <= 0 {
		error := Error{Message: "Notes could not be retrieved."}
		data, _ := json.Marshal(error)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
	}

	data, _ := json.Marshal(notes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// var note Note
	for _, item := range notes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Note{})
	error := Error{Message: "Cant find note by that id"}
	data, _ := json.Marshal(error)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNotFound)
	w.Write(data)
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
	error := Error{Message: "Can't delete a note that doesn't exist."}
	data, _ := json.Marshal(error)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNotFound)
	w.Write(data)
}

func EditNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var Note Note
	_ = json.NewDecoder(req.Body).Decode(&Note)
	for index, item := range notes {
		if item.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			notes = append(notes, Note)
			json.NewEncoder(w).Encode(notes)
			break
		}
	}
	error := Error{Message: "Cant edit note by that id"}
	data, _ := json.Marshal(error)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNotFound)
	w.Write(data)

}

func main() {
	router := mux.NewRouter()

	notes = append(notes, Note{ID: "0", Title: "Note Title", TextBody: "Note Body", Tags: []string{"tag1", "tag2"}})
	notes = append(notes, Note{ID: "1", Title: "Note Title 2", TextBody: "Note Body 2"})

	router.HandleFunc("/note/get/all", GetNotesEndpoint).Methods("GET")
	router.HandleFunc("/note/get/{id}", GetNoteEndpoint).Methods("GET")
	router.HandleFunc("/note/create", CreateNoteEndpoint).Methods("POST")
	router.HandleFunc("/note/delete/{id}", DeleteNoteEndpoint).Methods("DELETE")
	router.HandleFunc("/note/edit/{id}", EditNoteEndpoint).Methods("PUT")

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":12345", router))
}
