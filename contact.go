package main

import (
	"context"
	"log"
	"os"

	"github.com/sigma-firma/gmailAPI"
	ohsheet "github.com/sigma-firma/googlesheetsapi"
	"github.com/sigma-firma/inboxer"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/sheets/v4"
)

func sendConf(c *contactForm) {
	msg := &inboxer.Msg{
		To:      c.Email,
		Subject: "Welcome aboard the ship, captain",
		Body: "Click the following link to fill out our questionnaire:" +
			"https://forms.gle/tuC2nuh5EBsjNQCc8" +
			"<br /><br /><div style=\"font-size: 3em; font-weight: bold;\"><div style=\"color: red; display: inline;\">Σ</div>firma</div>",
		ImagePath: "",
		MimeType:  "",
	}
	err := bobbyEmail(msg)
	if err != nil {
		log.Println(err)
	}
}

func bobbyEmail(msg *inboxer.Msg) error {
	msg.From = "me"
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
func formatContactEmail(cf *contactForm) string {
	return "<div style=\"white-space: pre-wrap; font-size:2em;\">" +
		"Name: " + cf.FirstName + " " + cf.LastName +
		"\nPhone: " + cf.Phone +
		"\nEmail: " + cf.Email +
		"\nNews Letter Opt In: " + cf.NewsLetter + "</div>"
}
