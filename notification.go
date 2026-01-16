package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Notification(url, message string) {
	payload := map[string]string{"text": "Reminder: " + message}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error encoding JSON:", err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("Error sending Slack message:", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Println("Slack message sent successfully!")
	} else {
		log.Printf("Failed to send Slack message. Status: %s\n", response.Status)
	}
}
