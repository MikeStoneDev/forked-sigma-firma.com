package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sigma-firma/inboxer"
)

func root(w http.ResponseWriter, r *http.Request) {
	exeTmpl(w, r, nil, "main.html")
}
func contact(w http.ResponseWriter, r *http.Request) {
	cf, err := marshalContact(r)
	if err != nil {
		log.Println(err)
	}

	if cf.Email == "" || cf.FirstName == cf.LastName {
		ajaxResponse(w, map[string]string{"success": "false", "message": "invalid form data"})
		return
	}

	msg := &inboxer.Msg{
		To:        "leadership@sigma-firma.com",
		Subject:   "New Contact",
		Body:      formatContactEmail(cf),
		ImagePath: "public/media/owlen.png",
		MimeType:  "png",
	}
	err = bobbyEmail(msg)
	if err != nil {
		log.Println(err)
		ajaxResponse(w, map[string]string{"success": "false"})
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
		ajaxResponse(w, map[string]string{"success": "false"})
		return
	}
	ajaxResponse(w, map[string]string{"success": "true"})
}
