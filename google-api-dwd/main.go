package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	serviceAccountKeyJSONPath := "PATH/TO/YOUR_FILE"
	b, err := ioutil.ReadFile(serviceAccountKeyJSONPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	config.Subject = "user@example.com"

	client := &http.Client{
		Transport: &oauth2.Transport{
			Source: config.TokenSource(context.TODO()),
		},
	}

	srv, err := drive.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	list, err := srv.Files.List().Context(context.TODO()).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for i, file := range list.Files {
		fmt.Printf("[%d] %s", i, file.Name)
	}
}
