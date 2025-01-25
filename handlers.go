package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sigma-firma/gmailAPI"
	ohsheet "github.com/sigma-firma/googlesheetsapi"
	"github.com/sigma-firma/inboxer"
	gmail "google.golang.org/api/gmail/v1"
)

func root(w http.ResponseWriter, r *http.Request) {
	exeTmpl(w, r, nil, "main.html")
}
func contact(w http.ResponseWriter, r *http.Request) {
	cf, err := marshalContact(r)
	if err != nil {
		log.Println(err)
	}

	msg := &inboxer.Msg{
		From:    "me",
		To:      "leadership@sigma-firma.com",
		Subject: "New Contact",
		Body:    makeEmailBody(cf),
	}

	// Connect to the gmail API service.
	ctx := context.Background()
	srv := gmailAPI.ConnectToService(ctx, os.Getenv("HOME")+"/credentials", gmail.MailGoogleComScope)
	if msg.Send(srv) != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false"})
		return
	}
	ajaxResponse(w, map[string]string{"success": "true"})
	addRow([]interface{}{cf.LastName, cf.FirstName, cf.Phone, cf.Email, cf.NewsLetter, cf.Questionnaire, false})
}

func addRow(row []interface{}) {
	// Connect to the API
	sheet := &ohsheet.Access{
		Token:       os.Getenv("HOME") + "/credentials/sheets-go-quickstart.json",
		Credentials: os.Getenv("HOME") + "/credentials/credentials.json",
		Scopes:      []string{"https://www.googleapis.com/auth/spreadsheets"},
	}
	srv := sheet.Connect()

	spreadsheetId := "1cZVwQaY8LqsIUwzbCm_yG8tcR5RDog9jD1sHJtF9mSA"

	// Write to the sheet
	writeRange := "A2"
	res, err := sheet.Write(srv, spreadsheetId, writeRange, row)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	fmt.Println(res)
}
func makeEmailBody(cf *contactForm) string {
	return "Name: " + cf.FirstName + " " + cf.LastName + "\nPhone: " +
		cf.Phone + "\nEmail: " + cf.Email + "\nNews Letter Opt In: " +
		cf.NewsLetter + "\nQuestionnaire: " + cf.Questionnaire
}
