package utils

import (
	"fmt"
	"log"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type ExpoNotification struct {
	Token string            `json:"token"`
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Data  map[string]string `json:"data,omitempty"`
}

func SendExpoNotification(notification *ExpoNotification) (string, error) {
	pushToken, err := expo.NewExponentPushToken(notification.Token)
	if err != nil {
		return "", fmt.Errorf("invalid push token: %v", err)
	}
	client := expo.NewPushClient(nil)

	message := &expo.PushMessage{
		To:       []expo.ExponentPushToken{pushToken},
		Title:    notification.Title,
		Body:     notification.Body,
		Data:     notification.Data,
		Sound:    "default",
		Priority: expo.DefaultPriority,
	}

	response, err := client.Publish(message)
	if err != nil {
		return "", fmt.Errorf("failed to send push notification: %v", err)
	}
	if validateErr := response.ValidateResponse(); validateErr != nil {
		log.Printf("Warning: push notification response validation failed: %v", validateErr)
	}

	if len(response.ID) > 0 {
		return response.ID, nil
	}

	return "", fmt.Errorf("no push receipt ID received")
}
