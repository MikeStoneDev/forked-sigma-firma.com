package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/sigma-firma/gmailAPI"
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
}

func makeEmailBody(cf *contactForm) string {
	return "Name: " + cf.FirstName + " " + cf.LastName + "\nPhone: " +
		cf.Phone + "\nEmail: " + cf.Email + "\nNews Letter Opt In: " +
		cf.NewsLetter + "\nQuestionnaire: " + cf.Questionnaire
}
