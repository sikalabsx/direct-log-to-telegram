package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sikalabsx/direct-log-to-telegram/pkg/direct_log_to_telegram"
)

func main() {
	// Send a simple log message to Telegram
	err := direct_log_to_telegram.Log("Hello from direct-log-to-telegram!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message sent successfully!")

	// Send a more detailed log message
	err = direct_log_to_telegram.Log(fmt.Sprintf("Application started successfully at %s", time.Now().Format(time.RFC3339)))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Second message sent successfully!")
}
