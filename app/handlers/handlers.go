package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	port := ":" + os.Getenv("PORT")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", nil)

	log.Fatal(http.ListenAndServe(port, router))
}
