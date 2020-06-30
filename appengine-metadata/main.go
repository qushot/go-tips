package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
)

func main() {
	http.HandleFunc("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	str := ""

	numericProjectID, err := metadata.NumericProjectID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	str += fmt.Sprintf("NumericProjectID: %s, ", numericProjectID)

	projectID, err := metadata.ProjectID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	str += fmt.Sprintf("ProjectID: %s, ", projectID)

	zone, err := metadata.Zone()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	str += fmt.Sprintf("Zone: %s, ", zone)

	email, err := metadata.Email("default")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	str += fmt.Sprintf("Eamil: %s, ", email)

	scopes, err := metadata.Scopes("default")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	str += fmt.Sprintf("Scopes: %s, ", scopes)

	token, err := metadata.Get("/instance/service-accounts/default/token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	str += fmt.Sprintf("Token: %s, ", token)

	fmt.Fprintf(w, "%s", str)
}
