package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/go-lang-rest-api/pkg/notes"
)

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

	router.HandleFunc("/note/get/all", notes.GetAllNotesEndpoint).Methods("GET")
	router.HandleFunc("/note/get/{id}", notes.GetNoteEndpoint).Methods("GET")
	router.HandleFunc("/note/create", notes.CreateNoteEndpoint).Methods("POST")
	router.HandleFunc("/note/delete/{id}", notes.DeleteNoteEndpoint).Methods("DELETE")
	router.HandleFunc("/note/edit/{id}", notes.EditNoteEndpoint).Methods("PUT")

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":12345", router))
}
