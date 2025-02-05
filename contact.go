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

var confirmationBody string = "<div><a href=\"https://forms.gle/tuC2nuh5EBsjNQCc8\">" +
	"Click here to fill out our questionnaire.</a></div>" +
	"<div style=\"font-size: 3em; font-weight: bold;\"><div style=\"color: red; display: inline;\">Σ</div>firma</div>"

type contactForm struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NewsLetter    string `json:"news_letter"`
	Questionnaire string `json:"questionnaire"`
}

func contact(w http.ResponseWriter, r *http.Request) {
	cf, err := marshalContact(r)
	if err != nil {
		ajaxResponse(w, map[string]string{"success": "false", "error": "invalid form data"})
		log.Println(err)
		return
	}

	err = sendAll(cf)
	if err != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false", "error": "invalid form data"})
	}

	ajaxResponse(w, map[string]string{"success": "true"})
}

func sendAll(cf *contactForm) error {
	err := sendAlertEmail(cf)
	if err != nil {
		log.Println(err)
		return err
	}
	err = sendConf(cf)
	if err != nil {
		log.Println(err)
		return err
	}
	err = sendToSheet(cf)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
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
	if t.FirstName == "" || t.LastName == "" || t.Phone == "" || t.Email == "" {
		return t, errors.New("Invalid Input")
	}
	return t, nil
}
func sendToSheet(cf *contactForm) error {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return err
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
		return err
	}
	return nil
}
func sendAlertEmail(cf *contactForm) error {
	msg := &inboxer.Msg{
		To:        "leadership@sigma-firma.com",
		Subject:   "New Contact",
		Body:      formatContactEmail(cf),
		ImagePath: "",
		MimeType:  "",
	}
	return bobbyEmail(msg)
}
func sendConf(c *contactForm) error {
	msg := &inboxer.Msg{
		To:        c.Email,
		Subject:   "Welcome aboard the ship, captain",
		Body:      confirmationBody,
		ImagePath: "",
		MimeType:  "",
	}
	return bobbyEmail(msg)
}

func bobbyEmail(msg *inboxer.Msg) error {
	msg.From = "me"
	// srv := gmailAPI.ConnectToService(context.Background(), os.Getenv("HOME")+"/credentials", gmail.MailGoogleComScope)
	return msg.Send(connectToGoogleAPI())
}

func connectToGoogleAPI() *gmail.Service {
	return gmailAPI.ConnectToService(
		context.Background(),
		os.Getenv("HOME")+"/credentials",
		gmail.MailGoogleComScope,
	)
}

func autoRefreshGoogleToken() {
	for {
		call := connectToGoogleAPI().Users.GetProfile("me")

		_, err := call.Do()
		if err != nil {
			log.Println(err)
		}

		// refresh the token once every twelve hours
		time.Sleep(12 * time.Hour)
	}
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
