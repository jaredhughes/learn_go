package main

import (
	"net/http"

	"github.com/tsawler/toolbox"
)

func (app *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	_, err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestPayload.To,
	}

	tools.WriteJSON(w, http.StatusAccepted, payload)
}
