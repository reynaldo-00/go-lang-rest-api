package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Note structure
type Note struct {
	// ID of note in databse
	ID string `json:"id,omitempty"`
	// Title for note
	Title string `json:"title,omitempty"`
	// Body for the note
	TextBody string `json:"textBody,omitempty"`
	// Tags for note
	Tags []string `json:"tags"`
}

// Error structure
type Error struct {
	// Error mesasge
	Message string `json:"message,omitempty"`
}

var notes []Note

var noteID int

func main() {
	notes = append(
		notes,
		Note{
			ID:       "0",
			Title:    "Note Title",
			TextBody: "Note Body",
			Tags:     []string{"tag1", "tag2"},
		},
		Note{
			ID:       "1",
			Title:    "Note Title 2",
			TextBody: "Note Body 2",
		},
	)

	router := mux.NewRouter()

	router.HandleFunc("/note/get/all", GetAllNotesEndpoint).Methods("GET")
	router.HandleFunc("/note/get/{id}", GetNoteEndpoint).Methods("GET")
	router.HandleFunc("/note/create", CreateNoteEndpoint).Methods("POST")
	router.HandleFunc("/note/delete/{id}", DeleteNoteEndpoint).Methods("DELETE")
	router.HandleFunc("/note/edit/{id}", EditNoteEndpoint).Methods("PUT")

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":12345", router))
}

// GetAllNotesEndpoint gets all the notes from notes
func GetAllNotesEndpoint(w http.ResponseWriter, req *http.Request) {
	// json.NewEncoder(w).Encode(notes)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if len(notes) <= 0 {
		error := Error{Message: "Notes could not be retrieved."}
		data, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
	}

	data, _ := json.Marshal(notes)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetNoteEndpoint gets a specific note from notes
func GetNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(req)

	for _, item := range notes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	error := Error{Message: "Cant find note by that id"}
	data, _ := json.Marshal(error)
	w.WriteHeader(http.StatusNotFound)
	w.Write(data)
}

// CreateNoteEndpoint creates a new note in notes slice
func CreateNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	var n Note
	_ = json.NewDecoder(req.Body).Decode(&n)

	n.ID = string(len(notes))
	notes = append(notes, n)
	json.NewEncoder(w).Encode(n.ID)
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
