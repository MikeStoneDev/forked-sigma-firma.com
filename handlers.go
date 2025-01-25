package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sigma-firma/gmailAPI"
	ohsheet "github.com/sigma-firma/googlesheetsapi"
	"github.com/sigma-firma/inboxer"
	gmail "google.golang.org/api/gmail/v1"
	sheets "google.golang.org/api/sheets/v4"
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
		To:      "leadership@sigma-firma.com",
		Subject: "New Contact",
		Body:    makeEmailBody(cf),
	}
	err = bobbyEmail(msg)
	if err != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false"})
		return
	}

	_, err = addRow([]interface{}{cf.LastName, cf.FirstName, time.Now().Local().UTC().Format("Jan 02 2006"), cf.Phone, cf.Email, cf.NewsLetter, cf.Questionnaire, false})
	if err != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false"})
		return
	}
	ajaxResponse(w, map[string]string{"success": "true"})
}

func bobbyEmail(msg *inboxer.Msg) error {
	msg.From = "noreply@sigma-firma.com"
	ctx := context.Background()
	srv := gmailAPI.ConnectToService(ctx, os.Getenv("HOME")+"/credentials", gmail.MailGoogleComScope)
	return msg.Send(srv)
}

func addRow(row []interface{}) (*sheets.AppendValuesResponse, error) {
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
	return sheet.Write(srv, spreadsheetId, writeRange, row)
}
func makeEmailBody(cf *contactForm) string {
	return "Name: " + cf.FirstName + " " + cf.LastName + "\nPhone: " +
		cf.Phone + "\nEmail: " + cf.Email + "\nNews Letter Opt In: " +
		cf.NewsLetter + "\nQuestionnaire: " + cf.Questionnaire
}
