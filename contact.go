package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/sheets/v4"

	"github.com/sigma-firma/gsheet"
)

// var access *gsheet.Access = gsheet.NewAccess(
//
//	os.Getenv("HOME")+"/credentials/credentials.json",
//	os.Getenv("HOME")+"/credentials/token.json",
//	[]string{gmail.GmailComposeScope, sheets.SpreadsheetsScope},
//
// )
// var gm *gsheet.Gmailer = access.Gmail()
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
func sendAll(cf *contactForm) error {
	_, err := sendAlertEmail(cf)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = sendConf(cf)
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
func sendAlertEmail(cf *contactForm) (*gmail.Message, error) {
	msg := &gsheet.Msg{
		To:        "leadership@sigma-firma.com",
		Subject:   "New Contact",
		Body:      formatContactEmail(cf),
		ImagePath: "",
		MimeType:  "",
	}
	return bobbyEmail(msg)
}
func sendConf(c *contactForm) (*gmail.Message, error) {
	msg := &gsheet.Msg{
		To:        c.Email,
		Subject:   "Welcome aboard the ship, captain",
		Body:      confirmationBody,
		ImagePath: "",
		MimeType:  "",
	}
	return bobbyEmail(msg)
}

func bobbyEmail(msg *gsheet.Msg) (*gmail.Message, error) {
	msg.From = "me"
	return &gmail.Message{}, nil
	// gm.Service.Users.Messages.Send(msg.From, msg.Form()).Do()
}

func formatContactEmail(cf *contactForm) string {
	return "<div style=\"white-space: pre-wrap; font-size:2em;\">" +
		"Name: " + cf.FirstName + " " + cf.LastName +
		"\nPhone: " + cf.Phone +
		"\nEmail: " + cf.Email +
		"\nNews Letter Opt In: " + cf.NewsLetter + "</div>"
}

func addRow(row []interface{}) (*sheets.AppendValuesResponse, error) {
	// sh := access.Sheets()
	// var req *gsheet.SpreadSheet = &gsheet.SpreadSheet{
	// 	ID:               "1cZVwQaY8LqsIUwzbCm_yG8tcR5RDog9jD1sHJtF9mSA",
	// 	WriteRange:       "A2",
	// 	Vals:             row,
	// 	ValueInputOption: "RAW",
	// }

	// Write to the sheet
	// return sh.AppendRow(req)
	return &sheets.AppendValuesResponse{}, nil
}
