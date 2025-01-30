package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sigma-firma/gmailAPI"
	ohsheet "github.com/sigma-firma/googlesheetsapi"
	"github.com/sigma-firma/inboxer"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/sheets/v4"
)

type contactForm struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NewsLetter    string `json:"news_letter"`
	Questionnaire string `json:"questionnaire"`
}

// marshalCredentials is used convert a request body into a credentials{}
// struct
func marshalContact(r *http.Request) (*contactForm, error) {
	t := &contactForm{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(t)
	if err != nil {
		return t, err
	}
	log.Println(t.Phone)
	if t.FirstName == "" || t.LastName == "" || t.Phone == "" || t.Email == "" {
		return t, errors.New("Invalid Input")
	}
	return t, nil
}
func contact(w http.ResponseWriter, r *http.Request) {
	cf, err := marshalContact(r)
	if err != nil {
		ajaxResponse(w, map[string]string{"success": "false", "error": "invalid form data"})
		log.Println(err)
		return
	}

	if cf.Email == "" || cf.FirstName == cf.LastName {
		ajaxResponse(w, map[string]string{"success": "false", "error": "invalid form data"})
		return
	}

	msg := &inboxer.Msg{
		To:        "leadership@sigma-firma.com",
		Subject:   "New Contact",
		Body:      formatContactEmail(cf),
		ImagePath: "",
		MimeType:  "",
	}
	err = bobbyEmail(msg)
	if err != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false", "error": "invalid form data"})
		return
	}
	sendConf(cf)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)
	}
	_, err = addRow([]interface{}{
		cf.LastName,
		cf.FirstName,
		time.Now().UTC().In(loc).Format("Jan 02 2006 03:04:05 PM"),
		cf.Email,
		cf.Phone,
		cf.NewsLetter,
		cf.Questionnaire,
		false,
	})
	if err != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false", "error": "invalid form data"})
		return
	}
	ajaxResponse(w, map[string]string{"success": "true"})
}
func sendConf(c *contactForm) {
	msg := &inboxer.Msg{
		To:      c.Email,
		Subject: "Welcome aboard the ship, captain",
		Body: "<div><a href=\"https://forms.gle/tuC2nuh5EBsjNQCc8\">" +
			"Click here to fill out our questionnaire.</a></div>" +
			"<div style=\"font-size: 3em; font-weight: bold;\"><div style=\"color: red; display: inline;\">Σ</div>firma</div>",
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
