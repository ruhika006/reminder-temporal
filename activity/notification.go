package activity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Notification(url, message string) error{
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
		return nil
	} else {
		return fmt.Errorf("Failed to send Slack message Status: %s\n", response.Status)
	}
}
