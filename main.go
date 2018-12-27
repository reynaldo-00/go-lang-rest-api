package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	router.HandleFunc("/note/get/all", GetAllNotesEndpoint).Methods("GET")
	router.HandleFunc("/note/get/{id}", GetNoteEndpoint).Methods("GET")
	router.HandleFunc("/note/create", CreateNoteEndpoint).Methods("POST")
	router.HandleFunc("/note/delete/{id}", DeleteNoteEndpoint).Methods("DELETE")
	router.HandleFunc("/note/edit/{id}", EditNoteEndpoint).Methods("PUT")

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":12345", router))
}
