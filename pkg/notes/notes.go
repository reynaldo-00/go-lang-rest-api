package notes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var notes []Note

var noteID int

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

// Success message when req is succesful
type Success struct {
	Message string `json:"mesage,omitempty"`
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

// DeleteNoteEndpoint deletes note with matching id
func DeleteNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(req)
	deleted := 0

	for index, item := range notes {
		if item.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			deleted = deleted + 1
		}
	}

	if deleted == 0 {
		error := Error{Message: "Can't delete a note that doesn't exist."}
		data, _ := json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)
		w.Write(data)
	}
	s := Success{Message: "Deleted " + string(deleted) + "notes"}
	data, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

// EditNoteEndpoint edits note with matching id from request
func EditNoteEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(req)
	var Note Note
	_ = json.NewDecoder(req.Body).Decode(&Note)
	updated := 0

	for index, item := range notes {
		if item.ID == params["id"] {
			notes[index] = item
			json.NewEncoder(w).Encode(notes)
			updated = updated + 1
		}
	}

	if updated == 0 {
		error := Error{Message: "Can't delete a note that doesn't exist."}
		data, _ := json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)
		w.Write(data)
	}

	s := Success{Message: "Updated " + string(updated) + "notes"}
	data, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
